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
  student_name: '',
  course_name: '',
  status: null,
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
    title: '学员姓名',
    key: 'student_name',
    width: 80,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '课程名称',
    key: 'course_name',
    width: 100,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '教练姓名',
    key: 'coach_name',
    width: 80,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '预约状态',
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
    title: '预约时间',
    key: 'reservation_time',
    align: 'center',
    width: 150,
    render(row) {
      return h('span', formatDate(row.reservation_time, 'YYYY-MM-DD HH:mm'))
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
    title: '备注',
    key: 'remark',
    width: 120,
    align: 'center',
    ellipsis: { tooltip: true },
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
      if (row.status === 1) {
        actions.push(
          h(
            NButton,
            {
              size: 'small',
              type: 'success',
              onClick: () => handleUpdateStatus(row.id, 2),
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
              onClick: () => handleUpdateStatus(row.id, 3),
            },
            {
              default: () => '取消',
              icon: () => h('i', { class: 'i-material-symbols:close' }),
            },
          ),
        )
      }

      // 已确认状态可以标记为完成
      if (row.status === 2) {
        actions.push(
          h(
            NButton,
            {
              size: 'small',
              type: 'info',
              onClick: () => handleUpdateStatus(row.id, 4),
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
        <QueryItem label="学员姓名" :label-width="80">
          <NInput v-model:value="queryItems.student_name" clearable placeholder="请输入学员姓名" />
        </QueryItem>
        <QueryItem label="课程名称" :label-width="80">
          <NInput v-model:value="queryItems.course_name" clearable placeholder="请输入课程名称" />
        </QueryItem>
        <QueryItem label="预约状态" :label-width="80">
          <NSelect
            v-model:value="queryItems.status"
            clearable
            placeholder="请选择预约状态"
            :options="statusOptions"
          />
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
        <NFormItem label="学员姓名" path="student_name">
          <NInput v-model:value="modalForm.student_name" disabled />
        </NFormItem>
        <NFormItem label="课程名称" path="course_name">
          <NInput v-model:value="modalForm.course_name" disabled />
        </NFormItem>
        <NFormItem label="教练姓名" path="coach_name">
          <NInput v-model:value="modalForm.coach_name" disabled />
        </NFormItem>
        <NFormItem label="预约状态" path="status">
          <NSelect
            v-model:value="modalForm.status"
            :options="statusOptions"
          />
        </NFormItem>
        <NFormItem label="预约时间" path="reservation_time">
          <NInput v-model:value="modalForm.reservation_time" disabled />
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
