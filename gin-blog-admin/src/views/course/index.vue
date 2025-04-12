<script setup>
import { h, reactive, ref } from 'vue'
import { NButton, NPopconfirm, NSpace, useMessage } from 'naive-ui'
import { useCRUD } from '@/composables/useCRUD'
import api from '@/api'

const message = useMessage()

// 表格列配置
const columns = [
  { title: 'ID', key: 'id', width: 80 },
  { title: '课程名称', key: 'name' },
  { title: '课程描述', key: 'description' },
  { title: '教练', key: 'coach_name' },
  { title: '课程时长(分钟)', key: 'duration' },
  { title: '最大人数', key: 'max_students' },
  { title: '创建时间', key: 'created_at' },
  {
    title: '操作',
    key: 'actions',
    width: 200,
    render() {
      return h(NSpace, { justify: 'center' }, {
        default: () => [
          h(NButton, {
            size: 'small',
            type: 'primary',
            // onClick: () => handleEdit(row),
          }, { default: () => '编辑' }),
          h(NPopconfirm, {
            // onPositiveClick: () => handleDelete(row.id),
          }, {
            trigger: () => h(NButton, {
              size: 'small',
              type: 'error',
            }, { default: () => '删除' }),
            default: () => '确定删除？',
          }),
        ],
      })
    },
  },
]

// 表单字段配置
const formConfig = {
  name: {
    label: '课程名称',
    type: 'input',
    rule: { required: true, message: '请输入课程名称', trigger: 'blur' },
  },
  description: {
    label: '课程描述',
    type: 'textarea',
    rule: { required: true, message: '请输入课程描述', trigger: 'blur' },
  },
  coach_id: {
    label: '教练',
    type: 'select',
    rule: { required: true, message: '请选择教练', trigger: 'blur' },
    options: [],
  },
  duration: {
    label: '课程时长(分钟)',
    type: 'input-number',
    rule: { required: true, message: '请输入课程时长', trigger: 'blur' },
  },
  max_students: {
    label: '最大人数',
    type: 'input-number',
    rule: { required: true, message: '请输入最大人数', trigger: 'blur' },
  },
}

// 使用CRUD组合式API
const {
  tableRef,
  modalRef,
  formRef,
  pagination,
  modalType,
  formValue,
  loading,
  handleAdd,
  //   handleEdit,
  //   handleDelete,
  handleSubmit,
  fetchData,
} = useCRUD({
  name: '课程',
  columns,
  formConfig,
  api: {
    list: api.getCourseList,
    create: api.createCourse,
    update: api.updateCourse,
    delete: api.deleteCourse,
  },
})

// 获取教练列表
async function fetchCoachOptions() {
  try {
    const res = await api.getCoachList({ page: 1, page_size: 100 })
    formConfig.coach_id.options = res.data.list.map(item => ({
      label: item.name,
      value: item.id,
    }))
  }
  catch (error) {
    console.error('获取教练列表失败', error)
    message.error('获取教练列表失败')
  }
}

// 初始化数据
fetchData()
fetchCoachOptions()
</script>

<template>
  <div>
    <crud-table
      ref="tableRef"
      v-model:pagination="pagination"
      :columns="columns"
      :loading="loading"
      @fetch-data="fetchData"
      @add="handleAdd"
    />

    <crud-modal
      ref="modalRef"
      :type="modalType"
      :title="modalType === 'add' ? '新增课程' : '编辑课程'"
      @submit="handleSubmit"
    >
      <crud-form
        ref="formRef"
        v-model:form-value="formValue"
        :form-config="formConfig"
      />
    </crud-modal>
  </div>
</template>
