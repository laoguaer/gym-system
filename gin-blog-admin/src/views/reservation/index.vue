<script setup>
import { h, onMounted, ref } from 'vue'
import { NButton, NForm, NFormItem, NInput, NPopconfirm, NSelect, NSpace, NTag } from 'naive-ui'

import CommonPage from '@/components/common/CommonPage.vue'
import QueryItem from '@/components/crud/QueryItem.vue'
import CrudModal from '@/components/crud/CrudModal.vue'
import CrudTable from '@/components/crud/CrudTable.vue'

import { formatDate } from '@/utils'
import { useCRUD } from '@/composables'
import api from '@/api'

defineOptions({ name: '预约管理' })

const $table = ref(null)
const queryItems = ref({
  user_id: null,
  course_id: null,
  status: null,
  date: '',
})

const {
  modalVisible,
  modalLoading,
  handleEdit,
  modalForm,
  modalFormRef,
} = useCRUD({
  name: '预约',
  doUpdate: api.updateReservation,
  refresh: () => $table.value?.handleSearch(),
})

// 预约状态选项
const statusOptions = [
  { label: '待审核', value: 1 },
  { label: '已确认', value: 2 },
  { label: '已取消', value: 3 },
  { label: '已完成', value: 4 },
]

const statusMap = {
  1: { name: '待审核', tag: 'warning' },
  2: { name: '已确认', tag: 'success' },
  3: { name: '已取消', tag: 'error' },
  4: { name: '已完成', tag: 'info' },
}

onMounted(() => {
  $table.value?.handleSearch()
})

const columns = [
  {
    title: '预约ID',
    key: 'Id',
    width: 60,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '用户ID',
    key: 'UserID',
    width: 80,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '课程ID',
    key: 'CourseID',
    width: 80,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '教练ID',
    key: 'CoachID',
    width: 80,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '课程标题',
    key: 'CourseTitle',
    width: 120,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '预约状态',
    key: 'Status',
    width: 80,
    align: 'center',
    render(row) {
      return h(
        NTag,
        { type: statusMap[row.Status]?.tag },
        { default: () => statusMap[row.Status]?.name || '未知' },
      )
    },
  },
  {
    title: '开始时间',
    key: 'StartTime',
    align: 'center',
    width: 150,
    render(row) {
      return h('span', row.StartTime)
    },
  },
  {
    title: '结束时间',
    key: 'EndTime',
    align: 'center',
    width: 150,
    render(row) {
      return h('span', row.EndTime)
    },
  },
  {
    title: '创建时间',
    key: 'CreatedAt',
    align: 'center',
    width: 150,
    render(row) {
      return h(
        NButton,
        { size: 'small', type: 'text', ghost: true },
        {
          default: () => row.CreatedAt,
          icon: () => h('i', { class: 'i-mdi:update' }),
        },
      )
    },
  },
  {
    title: '操作',
    key: 'actions',
    width: 200,
    align: 'center',
    fixed: 'right',
    render(row) {
      const actions = []

      // 待审核状态可以确认或取消
      if (row.Status === 1) {
        actions.push(
          h(
            NButton,
            {
              size: 'small',
              type: 'success',
              onClick: () => handleUpdateStatus(row.Id, 2),
            },
            {
              default: () => '确认',
              icon: () => h('i', { class: 'i-material-symbols:check' }),
            },
          ),
          h(
            NButton,
            {
              size: 'small',
              type: 'error',
              style: 'margin-left: 10px;',
              onClick: () => handleUpdateStatus(row.Id, 3),
            },
            {
              default: () => '取消',
              icon: () => h('i', { class: 'i-material-symbols:close' }),
            },
          ),
        )
      }

      // 已确认状态可以标记为完成
      if (row.Status === 2) {
        actions.push(
          h(
            NButton,
            {
              size: 'small',
              type: 'info',
              onClick: () => handleUpdateStatus(row.Id, 4),
            },
            {
              default: () => '完成',
              icon: () => h('i', { class: 'i-material-symbols:done-all' }),
            },
          ),
        )
      }

      // 所有状态都可以查看详情
      actions.push(
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            style: actions.length > 0 ? 'margin-left: 10px;' : '',
            onClick: () => handleEdit(row),
          },
          {
            default: () => '详情',
            icon: () => h('i', { class: 'i-material-symbols:info-outline' }),
          },
        ),
      )

      return actions
    },
  },
]

// 更新预约状态
async function handleUpdateStatus(id, status) {
  if (!id) {
    return
  }
  try {
    await api.updateReservationStatus(id, status)
    $message?.success('操作成功')
    $table.value?.handleSearch()
  }
  catch (err) {
    console.error(err)
  }
}
</script>

<template>
  <CommonPage title="预约管理">
    <CrudTable
      ref="$table"
      v-model:query-items="queryItems"
      :columns="columns"
      :get-data="api.getReservations"
      :scroll-x="1200"
    >
      <template #queryBar>
        <QueryItem label="用户ID" :label-width="80">
          <NInput v-model:value="queryItems.user_id" clearable placeholder="请输入用户ID" />
        </QueryItem>
        <QueryItem label="课程ID" :label-width="80">
          <NInput v-model:value="queryItems.course_id" clearable placeholder="请输入课程ID" />
        </QueryItem>
        <QueryItem label="预约状态" :label-width="80">
          <NSelect
            v-model:value="queryItems.status"
            clearable
            placeholder="请选择预约状态"
            :options="statusOptions"
          />
        </QueryItem>
        <QueryItem label="预约日期" :label-width="80">
          <NInput v-model:value="queryItems.date" clearable placeholder="请输入预约日期" />
        </QueryItem>
      </template>
    </CrudTable>

    <CrudModal
      v-model:visible="modalVisible"
      :loading="modalLoading"
      title="预约详情"
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
        <NFormItem label="用户ID" path="UserID">
          <NInput v-model:value="modalForm.UserID" disabled />
        </NFormItem>
        <NFormItem label="课程ID" path="CourseID">
          <NInput v-model:value="modalForm.CourseID" disabled />
        </NFormItem>
        <NFormItem label="教练ID" path="CoachID">
          <NInput v-model:value="modalForm.CoachID" disabled />
        </NFormItem>
        <NFormItem label="课程标题" path="CourseTitle">
          <NInput v-model:value="modalForm.CourseTitle" disabled />
        </NFormItem>
        <NFormItem label="预约状态" path="Status">
          <NSelect
            v-model:value="modalForm.Status"
            :options="statusOptions"
          />
        </NFormItem>
        <NFormItem label="开始时间" path="StartTime">
          <NInput v-model:value="modalForm.StartTime" disabled />
        </NFormItem>
        <NFormItem label="结束时间" path="EndTime">
          <NInput v-model:value="modalForm.EndTime" disabled />
        </NFormItem>
      </NForm>
    </CrudModal>
  </CommonPage>
</template>
