
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      

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
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
          <el-table-column align="left" label="关联用户表的用户 ID" prop="userId" width="120" />
          <el-table-column align="left" label="关联持仓表的持仓 ID" prop="positionId" width="120" />
          <el-table-column align="left" label="关联订单表的订单 ID" prop="orderId" width="120" />
          <el-table-column align="left" label="股票id" prop="stockId" width="120" />
          <el-table-column align="left" label="股票名称" prop="stockName" width="120" />
          <el-table-column align="left" label="触发类型" prop="triggerType" width="120" />
          <el-table-column align="left" label="触发价格" prop="triggerPrice" width="120" />
          <el-table-column align="left" label="操作类型" prop="operationType" width="120" />
          <el-table-column align="left" label="数量" prop="quantity" width="120" />
          <el-table-column align="left" label="委托状态" prop="STATUS" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateContractEntrustFunc(scope.row)">变更</el-button>
            <el-button v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
            <el-form-item label="关联用户表的用户 ID:"  prop="userId" >
              <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入关联用户表的用户 ID" />
            </el-form-item>
            <el-form-item label="关联持仓表的持仓 ID:"  prop="positionId" >
              <el-input v-model.number="formData.positionId" :clearable="true" placeholder="请输入关联持仓表的持仓 ID" />
            </el-form-item>
            <el-form-item label="关联订单表的订单 ID:"  prop="orderId" >
              <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入关联订单表的订单 ID" />
            </el-form-item>
            <el-form-item label="股票id:"  prop="stockId" >
              <el-input v-model.number="formData.stockId" :clearable="true" placeholder="请输入股票id" />
            </el-form-item>
            <el-form-item label="股票名称:"  prop="stockName" >
              <el-input v-model="formData.stockName" :clearable="true"  placeholder="请输入股票名称" />
            </el-form-item>
            <el-form-item label="触发类型:"  prop="triggerType" >
              <el-input v-model.number="formData.triggerType" :clearable="true" placeholder="请输入触发类型" />
            </el-form-item>
            <el-form-item label="触发价格:"  prop="triggerPrice" >
              <el-input-number v-model="formData.triggerPrice"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="操作类型:"  prop="operationType" >
              <el-input v-model.number="formData.operationType" :clearable="true" placeholder="请输入操作类型" />
            </el-form-item>
            <el-form-item label="数量:"  prop="quantity" >
              <el-input-number v-model="formData.quantity"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="委托状态:"  prop="STATUS" >
              <el-input v-model.number="formData.STATUS" :clearable="true" placeholder="请输入委托状态" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="关联用户表的用户 ID">
                        {{ detailFrom.userId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="关联持仓表的持仓 ID">
                        {{ detailFrom.positionId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="关联订单表的订单 ID">
                        {{ detailFrom.orderId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="股票id">
                        {{ detailFrom.stockId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="股票名称">
                        {{ detailFrom.stockName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="触发类型">
                        {{ detailFrom.triggerType }}
                    </el-descriptions-item>
                    <el-descriptions-item label="触发价格">
                        {{ detailFrom.triggerPrice }}
                    </el-descriptions-item>
                    <el-descriptions-item label="操作类型">
                        {{ detailFrom.operationType }}
                    </el-descriptions-item>
                    <el-descriptions-item label="数量">
                        {{ detailFrom.quantity }}
                    </el-descriptions-item>
                    <el-descriptions-item label="委托状态">
                        {{ detailFrom.STATUS }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createContractEntrust,
  deleteContractEntrust,
  deleteContractEntrustByIds,
  updateContractEntrust,
  findContractEntrust,
  getContractEntrustList
} from '@/api/contract/contractEntrust'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'




defineOptions({
    name: 'ContractEntrust'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            userId: undefined,
            positionId: undefined,
            orderId: undefined,
            stockId: undefined,
            stockName: '',
            triggerType: undefined,
            triggerPrice: 0,
            operationType: undefined,
            quantity: 0,
            STATUS: undefined,
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
  const table = await getContractEntrustList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteContractEntrustFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteContractEntrustByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateContractEntrustFunc = async(row) => {
    const res = await findContractEntrust({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteContractEntrustFunc = async (row) => {
    const res = await deleteContractEntrust({ ID: row.ID })
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
        userId: undefined,
        positionId: undefined,
        orderId: undefined,
        stockId: undefined,
        stockName: '',
        triggerType: undefined,
        triggerPrice: 0,
        operationType: undefined,
        quantity: 0,
        STATUS: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createContractEntrust(formData.value)
                  break
                case 'update':
                  res = await updateContractEntrust(formData.value)
                  break
                default:
                  res = await createContractEntrust(formData.value)
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
  const res = await findContractEntrust({ ID: row.ID })
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