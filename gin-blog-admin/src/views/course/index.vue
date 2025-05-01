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
  title: '',
  coach_id: null,
  tag_id: null,
  is_single: null,
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

// 课程类型选项（单人/团体）
const singleOptions = [
  { label: '单人课程', value: 1 },
  { label: '团体课程', value: 0 },
]

const singleMap = {
  1: { name: '单人课程', tag: 'info' },
  0: { name: '团体课程', tag: 'success' },
}

// 教练选项
const coachOptions = ref([])
// 标签选项
const tagOptions = ref([])

onMounted(async () => {
  $table.value?.handleSearch()
  try {
    // 获取教练选项 - 使用page和page_size参数，与CrudTable组件保持一致
    const coachResp = await api.getCoaches({ page: 1, size: 10 })
    if (coachResp && coachResp.data) {
      coachOptions.value = coachResp.data.page_data.map(coach => ({
        label: coach.name,
        value: coach.id,
      }))
    }

    // 获取标签选项
    const tagResp = await api.getTagOption()
    if (tagResp && tagResp.data) {
      tagOptions.value = tagResp.data
    }
  }
  catch (err) {
    console.error(err)
  }
})

const columns = [
  {
    title: '课程名称',
    key: 'title',
    width: 120,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '课程描述',
    key: 'description',
    width: 150,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '课程类型',
    key: 'is_single',
    width: 80,
    align: 'center',
    render(row) {
      return h(
        NTag,
        { type: singleMap[row.is_single]?.tag },
        { default: () => singleMap[row.is_single]?.name || '未知' },
      )
    },
  },
  {
    title: '标签',
    key: 'tags',
    width: 100,
    align: 'center',
    ellipsis: { tooltip: true },
    render(row) {
      if (!row.tag_list || row.tag_list.length === 0)
        return h('span', '无')
      return h(
        NSpace,
        { size: 'small', wrap: true },
        {
          default: () => row.tag_list.map(tag => h(
            NTag,
            { size: 'small', type: 'info' },
            { default: () => tag },
          )),
        },
      )
    },
  },
  {
    title: '教练',
    key: 'coach_name',
    width: 80,
    align: 'center',
    render(row) {
      return h('span', row.coach_name)
    },
  },
  {
    title: '最大容量',
    key: 'max_capacity',
    width: 80,
    align: 'center',
  },
  {
    title: '开始时间',
    key: 'start_time',
    width: 150,
    align: 'center',
    render(row) {
      return h('span', formatDate(row.start_time, 'YYYY-MM-DD HH:mm'))
    },
  },
  {
    title: '结束时间',
    key: 'end_time',
    width: 150,
    align: 'center',
    render(row) {
      return h('span', formatDate(row.end_time, 'YYYY-MM-DD HH:mm'))
    },
  },
  {
    title: '创建时间',
    key: 'created_at',
    align: 'center',
    width: 150,
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
          <NInput v-model:value="queryItems.title" clearable placeholder="请输入课程名称" />
        </QueryItem>
        <QueryItem label="教练" :label-width="80">
          <NSelect
            v-model:value="queryItems.coach_id"
            clearable
            placeholder="请选择教练"
            :options="coachOptions"
          />
        </QueryItem>
        <QueryItem label="标签" :label-width="80">
          <NSelect
            v-model:value="queryItems.tag_id"
            clearable
            placeholder="请选择标签"
            :options="tagOptions"
          />
        </QueryItem>
        <QueryItem label="课程类型" :label-width="80">
          <NSelect
            v-model:value="queryItems.is_single"
            clearable
            placeholder="请选择课程类型"
            :options="singleOptions"
          />
        </QueryItem>
      </template>

      <template #toolbar>
        <NButton type="primary" @click="handleEdit({})">
          <template #icon>
            <i class="i-material-symbols:add" />
          </template>
          新增课程
        </NButton>
      </template>
    </CrudTable>

    <CrudModal
      v-model:visible="modalVisible"
      :loading="modalLoading"
      :title="`${modalForm.id ? '编辑' : '新增'}课程`"
      @save="handleSave"
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
        <NFormItem label="课程名称" path="title" rule="required">
          <NInput v-model:value="modalForm.title" placeholder="请输入课程名称" />
        </NFormItem>
        <NFormItem label="课程描述" path="description">
          <NInput
            v-model:value="modalForm.description"
            type="textarea"
            placeholder="请输入课程描述"
          />
        </NFormItem>
        <NFormItem label="课程标签" path="tags">
          <NInput v-model:value="modalForm.tags" placeholder="请输入课程标签，多个标签用逗号分隔" />
        </NFormItem>
        <NFormItem label="开始时间" path="start_time" rule="required">
          <NTimePicker
            v-model:value="modalForm.start_time"
            type="datetime"
            clearable
            placeholder="请选择开始时间"
          />
        </NFormItem>
        <NFormItem label="结束时间" path="end_time" rule="required">
          <NTimePicker
            v-model:value="modalForm.end_time"
            type="datetime"
            clearable
            placeholder="请选择结束时间"
          />
        </NFormItem>
        <NFormItem label="教练" path="coach_id" rule="required">
          <NSelect
            v-model:value="modalForm.coach_id"
            placeholder="请选择教练"
            :options="coachOptions"
          />
        </NFormItem>
        <NFormItem label="最大容量" path="max_capacity" rule="required">
          <NInputNumber
            v-model:value="modalForm.max_capacity"
            placeholder="请输入最大容量"
            :min="1"
            :precision="0"
          />
        </NFormItem>
        <NFormItem label="课程类型" path="is_single" rule="required">
          <NSelect
            v-model:value="modalForm.is_single"
            placeholder="请选择课程类型"
            :options="singleOptions"
          />
        </NFormItem>
      </NForm>
    </CrudModal>
  </CommonPage>
</template>
