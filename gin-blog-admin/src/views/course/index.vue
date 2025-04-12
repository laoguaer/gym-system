<script setup>
import { h, onMounted, ref } from 'vue'
import { NButton, NForm, NFormItem, NInput, NInputNumber, NSelect, NSpace, NSwitch, NTag, NTimePicker } from 'naive-ui'

import CommonPage from '@/components/common/CommonPage.vue'
import QueryItem from '@/components/crud/QueryItem.vue'
import CrudModal from '@/components/crud/CrudModal.vue'
import CrudTable from '@/components/crud/CrudTable.vue'

import { formatDate } from '@/utils'
import { useCRUD } from '@/composables'
import api from '@/api'

defineOptions({ name: '课程列表' })

const $table = ref(null)
const queryItems = ref({
  name: '',
  type: null,
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
  name: '课程',
  doCreate: api.createCourse,
  doUpdate: api.updateCourse,
  refresh: () => $table.value?.handleSearch(),
})

// 课程类型选项
const typeOptions = [
  { label: '团体课', value: 1 },
  { label: '私教课', value: 2 },
  { label: '特色课', value: 3 },
]

const typeMap = {
  1: { name: '团体课', tag: 'info' },
  2: { name: '私教课', tag: 'success' },
  3: { name: '特色课', tag: 'warning' },
}

// 课程状态选项
const statusOptions = [
  { label: '正常', value: 1 },
  { label: '已满', value: 2 },
  { label: '停课', value: 3 },
]

const statusMap = {
  1: { name: '正常', tag: 'success' },
  2: { name: '已满', tag: 'warning' },
  3: { name: '停课', tag: 'error' },
}

onMounted(() => {
  $table.value?.handleSearch()
})

const columns = [
  {
    title: '课程名称',
    key: 'name',
    width: 100,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '课程类型',
    key: 'type',
    width: 80,
    align: 'center',
    render(row) {
      return h(
        NTag,
        { type: typeMap[row.type]?.tag },
        { default: () => typeMap[row.type]?.name || '未知' },
      )
    },
  },
  {
    title: '课程状态',
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
    title: '课程价格',
    key: 'price',
    width: 80,
    align: 'center',
    render(row) {
      return h('span', `¥${row.price}`)
    },
  },
  {
    title: '课程容量',
    key: 'capacity',
    width: 80,
    align: 'center',
  },
  {
    title: '课程时长',
    key: 'duration',
    width: 80,
    align: 'center',
    render(row) {
      return h('span', `${row.duration}分钟`)
    },
  },
  {
    title: '创建时间',
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
    title: '操作',
    key: 'actions',
    width: 120,
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
        h(
          NButton,
          {
            size: 'small',
            type: 'error',
            style: 'margin-left: 15px;',
            onClick: () => handleDelete([row.id]),
          },
          {
            default: () => '删除',
            icon: () => h('i', { class: 'i-material-symbols:delete-outline' }),
          },
        ),
      ]
    },
  },
]

// 删除课程
async function handleDelete(ids) {
  if (!ids || ids.length === 0) {
    return
  }
  try {
    await api.deleteCourse(ids)
    $message?.success('删除成功')
    $table.value?.handleSearch()
  }
  catch (err) {
    console.error(err)
  }
}
</script>

<template>
  <CommonPage title="课程列表">
    <CrudTable
      ref="$table"
      v-model:query-items="queryItems"
      :columns="columns"
      :get-data="api.getCourseList"
      :scroll-x="1200"
    >
      <template #queryBar>
        <QueryItem label="课程名称" :label-width="80">
          <NInput v-model:value="queryItems.name" clearable placeholder="请输入课程名称" />
        </QueryItem>
        <QueryItem label="课程类型" :label-width="80">
          <NSelect
            v-model:value="queryItems.type"
            clearable
            placeholder="请选择课程类型"
            :options="typeOptions"
          />
        </QueryItem>
        <QueryItem label="课程状态" :label-width="80">
          <NSelect
            v-model:value="queryItems.status"
            clearable
            placeholder="请选择课程状态"
            :options="statusOptions"
          />
        </QueryItem>
      </template>

      <template #toolbar>
        <NButton
          type="primary"
          @click="handleSave()"
        >
          新增课程
        </NButton>
      </template>
    </CrudTable>

    <CrudModal
      v-model:visible="modalVisible"
      :loading="modalLoading"
      :title="`${modalForm.id ? '编辑' : '新增'}课程`"
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
        <NFormItem label="课程名称" path="name" rule="required">
          <NInput v-model:value="modalForm.name" placeholder="请输入课程名称" />
        </NFormItem>
        <NFormItem label="课程类型" path="type" rule="required">
          <NSelect
            v-model:value="modalForm.type"
            placeholder="请选择课程类型"
            :options="typeOptions"
          />
        </NFormItem>
        <NFormItem label="课程状态" path="status" rule="required">
          <NSelect
            v-model:value="modalForm.status"
            placeholder="请选择课程状态"
            :options="statusOptions"
          />
        </NFormItem>
        <NFormItem label="课程价格" path="price" rule="required">
          <NInputNumber
            v-model:value="modalForm.price"
            placeholder="请输入课程价格"
            :min="0"
            :precision="2"
          />
        </NFormItem>
        <NFormItem label="课程容量" path="capacity" rule="required">
          <NInputNumber
            v-model:value="modalForm.capacity"
            placeholder="请输入课程容量"
            :min="1"
            :precision="0"
          />
        </NFormItem>
        <NFormItem label="课程时长(分钟)" path="duration" rule="required">
          <NInputNumber
            v-model:value="modalForm.duration"
            placeholder="请输入课程时长"
            :min="1"
            :precision="0"
          />
        </NFormItem>
        <NFormItem label="课程描述" path="description">
          <NInput
            v-model:value="modalForm.description"
            type="textarea"
            placeholder="请输入课程描述"
          />
        </NFormItem>
      </NForm>
    </CrudModal>
  </CommonPage>
</template>
