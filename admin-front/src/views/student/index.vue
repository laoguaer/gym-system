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

defineOptions({ name: '学员管理' })

const $table = ref(null)
const queryItems = ref({
  name: '',
  nickname: '',
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
    title: '用户名',
    key: 'username',
    width: 80,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '昵称',
    key: 'nickname',
    width: 80,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '邮箱',
    key: 'email',
    width: 100,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '简介',
    key: 'intro',
    width: 120,
    align: 'center',
    ellipsis: { tooltip: true },
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
    title: '更新时间',
    key: 'updated_at',
    align: 'center',
    width: 120,
    render(row) {
      return h(
        NButton,
        { size: 'small', type: 'text', ghost: true },
        {
          default: () => formatDate(row.updated_at),
          icon: () => h('i', { class: 'i-mdi:update' }),
        },
      )
    },
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
    // const resp = await api.getStudentCourses(id)
    // 这里可以实现查看学员课程记录的逻辑
    $message?.info('查看课程记录功能开发中')
  }
  catch (err) {
    console.error(err)
  }
}
</script>

<template>
  <CommonPage title="学员管理">
    <CrudTable
      ref="$table"
      v-model:query-items="queryItems"
      :columns="columns"
      :get-data="api.getStudents"
      :scroll-x="1200"
    >
      <template #queryBar>
        <QueryItem label="姓名" :label-width="80">
          <NInput v-model:value="queryItems.name" clearable placeholder="请输入姓名" />
        </QueryItem>
        <QueryItem label="昵称" :label-width="80">
          <NInput v-model:value="queryItems.nickname" clearable placeholder="请输入昵称" />
        </QueryItem>
      </template>
    </CrudTable>

    <CrudModal
      v-model:visible="modalVisible"
      :loading="modalLoading"
      title="学员详情"
      @save="handleEdit"
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
        <NFormItem label="用户名" path="username">
          <NInput v-model:value="modalForm.username" disabled />
        </NFormItem>
        <NFormItem label="邮箱" path="email">
          <NInput v-model:value="modalForm.email" disabled />
        </NFormItem>
        <NFormItem label="昵称" path="nickname">
          <NInput v-model:value="modalForm.nickname" disabled />
        </NFormItem>
        <NFormItem label="头像" path="avatar">
          <NInput v-model:value="modalForm.avatar" disabled />
        </NFormItem>
        <NFormItem label="简介" path="intro">
          <NInput
            v-model:value="modalForm.intro"
            type="textarea"
            placeholder="请输入简介信息"
          />
        </NFormItem>
      </NForm>
    </CrudModal>
  </CommonPage>
</template>
