
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
          <el-table-column align="left" label="id字段" prop="id" width="120" />
          <el-table-column align="left" label="langTag字段" prop="langTag" width="120" />
          <el-table-column align="left" label="messageId字段" prop="messageId" width="120" />
          <el-table-column align="left" label="templateData字段" prop="templateData" width="120" />
          <el-table-column align="left" label="categoryId字段" prop="categoryId" width="120" />
          <el-table-column align="left" label="错误码" prop="errorCode" width="120" />
         <el-table-column align="left" label="createdAt字段" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
         </el-table-column>
         <el-table-column align="left" label="updatedAt字段" prop="updatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.updatedAt) }}</template>
         </el-table-column>
         <el-table-column align="left" label="deletedAt字段" prop="deletedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.deletedAt) }}</template>
         </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateSysI18nLocalizeConfigFunc(scope.row)">变更</el-button>
            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="id字段:"  prop="id" >
              <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
            </el-form-item>
            <el-form-item label="langTag字段:"  prop="langTag" >
              <el-input v-model="formData.langTag" :clearable="true"  placeholder="请输入langTag字段" />
            </el-form-item>
            <el-form-item label="messageId字段:"  prop="messageId" >
              <el-input v-model="formData.messageId" :clearable="true"  placeholder="请输入messageId字段" />
            </el-form-item>
            <el-form-item label="templateData字段:"  prop="templateData" >
              <el-input v-model="formData.templateData" :clearable="true"  placeholder="请输入templateData字段" />
            </el-form-item>
            <el-form-item label="categoryId字段:"  prop="categoryId" >
              <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="请输入categoryId字段" />
            </el-form-item>
            <el-form-item label="错误码:"  prop="errorCode" >
              <el-input v-model.number="formData.errorCode" :clearable="true" placeholder="请输入错误码" />
            </el-form-item>
            <el-form-item label="createdAt字段:"  prop="createdAt" >
              <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="updatedAt字段:"  prop="updatedAt" >
              <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="deletedAt字段:"  prop="deletedAt" >
              <el-date-picker v-model="formData.deletedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="id字段">
                        {{ detailFrom.id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="langTag字段">
                        {{ detailFrom.langTag }}
                    </el-descriptions-item>
                    <el-descriptions-item label="messageId字段">
                        {{ detailFrom.messageId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="templateData字段">
                        {{ detailFrom.templateData }}
                    </el-descriptions-item>
                    <el-descriptions-item label="categoryId字段">
                        {{ detailFrom.categoryId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="错误码">
                        {{ detailFrom.errorCode }}
                    </el-descriptions-item>
                    <el-descriptions-item label="createdAt字段">
                        {{ detailFrom.createdAt }}
                    </el-descriptions-item>
                    <el-descriptions-item label="updatedAt字段">
                        {{ detailFrom.updatedAt }}
                    </el-descriptions-item>
                    <el-descriptions-item label="deletedAt字段">
                        {{ detailFrom.deletedAt }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createSysI18nLocalizeConfig,
  deleteSysI18nLocalizeConfig,
  deleteSysI18nLocalizeConfigByIds,
  updateSysI18nLocalizeConfig,
  findSysI18nLocalizeConfig,
  getSysI18nLocalizeConfigList
} from '@/api/i18n/sysI18nLocalizeConfig'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'SysI18nLocalizeConfig'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: undefined,
            langTag: '',
            messageId: '',
            templateData: '',
            categoryId: undefined,
            errorCode: undefined,
            createdAt: new Date(),
            updatedAt: new Date(),
            deletedAt: new Date(),
        })



// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getSysI18nLocalizeConfigList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteSysI18nLocalizeConfigFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteSysI18nLocalizeConfigByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateSysI18nLocalizeConfigFunc = async(row) => {
    const res = await findSysI18nLocalizeConfig({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteSysI18nLocalizeConfigFunc = async (row) => {
    const res = await deleteSysI18nLocalizeConfig({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        id: undefined,
        langTag: '',
        messageId: '',
        templateData: '',
        categoryId: undefined,
        errorCode: undefined,
        createdAt: new Date(),
        updatedAt: new Date(),
        deletedAt: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createSysI18nLocalizeConfig(formData.value)
                  break
                case 'update':
                  res = await updateSysI18nLocalizeConfig(formData.value)
                  break
                default:
                  res = await createSysI18nLocalizeConfig(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findSysI18nLocalizeConfig({ id: row.id })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>