<script setup>
import { h, onMounted, ref } from 'vue'
import { NButton, NCheckbox, NCheckboxGroup, NForm, NFormItem, NImage, NInput, NSelect, NSpace, NSwitch, NTag } from 'naive-ui'

import CommonPage from '@/components/common/CommonPage.vue'
import QueryItem from '@/components/crud/QueryItem.vue'
import CrudModal from '@/components/crud/CrudModal.vue'
import CrudTable from '@/components/crud/CrudTable.vue'

import { convertImgUrl, formatDate } from '@/utils'
import { useCRUD } from '@/composables'
import api from '@/api'

defineOptions({ name: '教练列表' })

const $table = ref(null)
const queryItems = ref({
  name: '',
  specialty: '',
  status: null,
})

const {
  modalVisible,
  modalLoading,
  handleSave,
  handleEdit,
  modalForm,
  modalFormRef,
} = useCRUD({
  name: '教练',
  doUpdate: api.updateCoach,
  refresh: () => $table.value?.handleSearch(),
})

// 教练状态选项
const statusOptions = [
  { label: '在职', value: 1 },
  { label: '离职', value: 2 },
  { label: '休假', value: 3 },
]

const statusMap = {
  1: { name: '在职', tag: 'success' },
  2: { name: '离职', tag: 'error' },
  3: { name: '休假', tag: 'warning' },
}

onMounted(() => {
  $table.value?.handleSearch()
})

const columns = [
  {
    title: '头像',
    key: 'avatar',
    width: 30,
    align: 'center',
    render(row) {
      return h(NImage, {
        'height': 30,
        'imgProps': { style: { 'border-radius': '3px' } },
        'src': convertImgUrl(row.avatar),
        'fallback-src': 'http://dummyimage.com/400x400', // 加载失败
        'show-toolbar-tooltip': true,
      })
    },
  },
  {
    title: '姓名',
    key: 'name',
    width: 60,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '性别',
    key: 'gender',
    width: 40,
    align: 'center',
    render(row) {
      return h('span', row.gender === 1 ? '男' : '女')
    },
  },
  {
    title: '专长',
    key: 'specialty',
    width: 80,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '电话',
    key: 'phone',
    width: 70,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '状态',
    key: 'status',
    width: 40,
    align: 'center',
    render(row) {
      return h(
        NTag,
        { type: statusMap[row.status]?.tag },
        { default: () => statusMap[row.status]?.name || '未知' },
      )
    },
  },
  {
    title: '简介',
    key: 'description',
    width: 100,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '创建时间',
    key: 'created_at',
    align: 'center',
    width: 70,
    render(row) {
      return h(
        NButton,
        { size: 'small', type: 'text', ghost: true },
        {
          default: () => formatDate(row.created_at),
          icon: () => h('i', { class: 'i-mdi:update' }),
        },
      )
    },
  },
  {
    title: '禁用',
    key: 'is_disable',
    width: 30,
    align: 'center',
    fixed: 'left',
    render(row) {
      return h(NSwitch, {
        size: 'small',
        rubberBand: false,
        value: row.is_disable,
        loading: !!row.publishing,
        onUpdateValue: () => handleUpdateDisable(row),
      })
    },
  },
  {
    title: '操作',
    key: 'actions',
    width: 60,
    align: 'center',
    fixed: 'right',
    render(row) {
      return [
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            onClick: () => handleEdit(row),
          },
          {
            default: () => '编辑',
            icon: () => h('i', { class: 'i-material-symbols:edit-outline' }),
          },
        ),
      ]
    },
  },
]

// 修改教练禁用状态
async function handleUpdateDisable(row) {
  if (!row.id) {
    return
  }
  row.publishing = true
  row.is_disable = !row.is_disable
  try {
    await api.updateCoachDisable(row.id, row.is_disable)
    $message?.success(row.is_disable ? '已禁用该教练' : '已取消禁用该教练')
    $table.value?.handleSearch()
  }
  catch (err) {
    row.is_disable = !row.is_disable
    console.error(err)
  }
  finally {
    row.publishing = false
  }
}
</script>

<template>
  <CommonPage title="教练列表">
    <CrudTable
      ref="$table"
      v-model:query-items="queryItems"
      :columns="columns"
      :get-data="api.getCoaches"
    >
      <template #queryBar>
        <QueryItem label="姓名" :label-width="40" :content-width="160">
          <NInput
            v-model:value="queryItems.name"
            clearable
            type="text"
            placeholder="请输入姓名"
            @keydown.enter="$table?.handleSearch()"
          />
        </QueryItem>
        <QueryItem label="专长" :label-width="40" :content-width="160">
          <NInput
            v-model:value="queryItems.specialty"
            clearable
            type="text"
            placeholder="请输入专长"
            @keydown.enter="$table?.handleSearch()"
          />
        </QueryItem>
        <QueryItem label="状态" :label-width="40" :content-width="160">
          <NSelect
            v-model:value="queryItems.status"
            clearable
            filterable
            placeholder="请选择状态"
            :options="statusOptions"
            @update:value="$table?.handleSearch()"
          />
        </QueryItem>
      </template>
    </CrudTable>

    <CrudModal
      v-model:visible="modalVisible"
      title="修改教练"
      :loading="modalLoading"
      @save="handleSave"
    >
      <NForm
        ref="modalFormRef"
        label-placement="left"
        label-align="left"
        :label-width="80"
        :model="modalForm"
      >
        <NFormItem label="姓名" path="name">
          <NInput
            v-model:value="modalForm.name"
            clearable
            placeholder="请输入姓名"
          />
        </NFormItem>
        <NFormItem label="性别" path="gender">
          <NSelect
            v-model:value="modalForm.gender"
            :options="[{ label: '男', value: 1 }, { label: '女', value: 2 }]"
            placeholder="请选择性别"
          />
        </NFormItem>
        <NFormItem label="专长" path="specialty">
          <NInput
            v-model:value="modalForm.specialty"
            clearable
            placeholder="请输入专长"
          />
        </NFormItem>
        <NFormItem label="电话" path="phone">
          <NInput
            v-model:value="modalForm.phone"
            clearable
            placeholder="请输入电话"
          />
        </NFormItem>
        <NFormItem label="状态" path="status">
          <NSelect
            v-model:value="modalForm.status"
            :options="statusOptions"
            placeholder="请选择状态"
          />
        </NFormItem>
        <NFormItem label="简介" path="description">
          <NInput
            v-model:value="modalForm.description"
            type="textarea"
            clearable
            placeholder="请输入简介"
          />
        </NFormItem>
      </NForm>
    </CrudModal>
  </CommonPage>
</template>
