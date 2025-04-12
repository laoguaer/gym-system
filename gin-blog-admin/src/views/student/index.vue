<script setup>
import { h, onMounted, ref } from 'vue'
import { NButton, NForm, NFormItem, NImage, NInput, NSelect, NSpace, NTag } from 'naive-ui'

import CommonPage from '@/components/common/CommonPage.vue'
import QueryItem from '@/components/crud/QueryItem.vue'
import CrudModal from '@/components/crud/CrudModal.vue'
import CrudTable from '@/components/crud/CrudTable.vue'

import { convertImgUrl, formatDate } from '@/utils'
import { useCRUD } from '@/composables'
import api from '@/api'

defineOptions({ name: '我的学员' })

const $table = ref(null)
const queryItems = ref({
  name: '',
  phone: '',
  status: null,
})

const {
  modalVisible,
  modalLoading,
  handleEdit,
  modalForm,
  modalFormRef,
} = useCRUD({
  name: '学员',
  doUpdate: api.updateStudent,
  refresh: () => $table.value?.handleSearch(),
})

// 学员状态选项
const statusOptions = [
  { label: '正常', value: 1 },
  { label: '已过期', value: 2 },
  { label: '已暂停', value: 3 },
]

const statusMap = {
  1: { name: '正常', tag: 'success' },
  2: { name: '已过期', tag: 'error' },
  3: { name: '已暂停', tag: 'warning' },
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
    width: 80,
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
    title: '电话',
    key: 'phone',
    width: 100,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '会员状态',
    key: 'status',
    width: 80,
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
    title: '会员到期时间',
    key: 'expire_time',
    align: 'center',
    width: 120,
    render(row) {
      return h('span', formatDate(row.expire_time, 'YYYY-MM-DD'))
    },
  },
  {
    title: '注册时间',
    key: 'created_at',
    align: 'center',
    width: 120,
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
    title: '备注',
    key: 'remark',
    width: 120,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '操作',
    key: 'actions',
    width: 180,
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
            default: () => '详情',
            icon: () => h('i', { class: 'i-material-symbols:info-outline' }),
          },
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'success',
            style: 'margin-left: 10px;',
            onClick: () => handleViewCourses(row.id),
          },
          {
            default: () => '课程记录',
            icon: () => h('i', { class: 'i-material-symbols:history' }),
          },
        ),
      ]
    },
  },
]

// 查看学员课程记录
async function handleViewCourses(id) {
  if (!id) {
    return
  }
  try {
    const resp = await api.getStudentCourses(id)
    // 这里可以实现查看学员课程记录的逻辑
    $message?.info('查看课程记录功能开发中')
  }
  catch (err) {
    console.error(err)
  }
}
</script>

<template>
  <CommonPage title="我的学员">
    <CrudTable
      ref="$table"
      v-model:query-items="queryItems"
      :columns="columns"
      :get-data="api.getStudents"
      :scroll-x="1200"
    >
      <template #queryBar>
        <QueryItem label="学员姓名" :label-width="80">
          <NInput v-model:value="queryItems.name" clearable placeholder="请输入学员姓名" />
        </QueryItem>
        <QueryItem label="电话" :label-width="80">
          <NInput v-model:value="queryItems.phone" clearable placeholder="请输入电话号码" />
        </QueryItem>
        <QueryItem label="会员状态" :label-width="80">
          <NSelect
            v-model:value="queryItems.status"
            clearable
            placeholder="请选择会员状态"
            :options="statusOptions"
          />
        </QueryItem>
      </template>
    </CrudTable>

    <CrudModal
      v-model:visible="modalVisible"
      :loading="modalLoading"
      title="学员详情"
      @submit="modalFormRef?.submit()"
    >
      <NForm
        ref="modalFormRef"
        :model="modalForm"
        label-placement="left"
        label-width="auto"
        require-mark-placement="right-hanging"
        :style="{
          maxWidth: '640px',
        }"
      >
        <NFormItem label="学员姓名" path="name">
          <NInput v-model:value="modalForm.name" disabled />
        </NFormItem>
        <NFormItem label="性别" path="gender">
          <NSelect
            v-model:value="modalForm.gender"
            :options="[{ label: '男', value: 1 }, { label: '女', value: 2 }]"
            disabled
          />
        </NFormItem>
        <NFormItem label="电话" path="phone">
          <NInput v-model:value="modalForm.phone" disabled />
        </NFormItem>
        <NFormItem label="会员状态" path="status">
          <NSelect
            v-model:value="modalForm.status"
            :options="statusOptions"
          />
        </NFormItem>
        <NFormItem label="会员到期时间" path="expire_time">
          <NInput v-model:value="modalForm.expire_time" disabled />
        </NFormItem>
        <NFormItem label="备注" path="remark">
          <NInput
            v-model:value="modalForm.remark"
            type="textarea"
            placeholder="请输入备注信息"
          />
        </NFormItem>
      </NForm>
    </CrudModal>
  </CommonPage>
</template>
