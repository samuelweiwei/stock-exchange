
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
          <el-table-column align="left" label="用户id" prop="uid" width="120" />
          <el-table-column align="left" label="理财产品ID" prop="productId" width="120" />
          <el-table-column align="left" label="下单ID" prop="subscribeId" width="120" />
          <el-table-column align="left" label="当前收益" prop="earnings" width="120" />
          <el-table-column align="left" label="当天利率" prop="interestRates" width="120" />
          <el-table-column align="left" label="买入数量" prop="boughtNum" width="120" />
        <el-table-column align="left" label="模拟|||默认不是模拟" prop="isMoni" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isMoni) }}</template>
        </el-table-column>
          <el-table-column align="left" label="到帐时间" prop="offeredAt" width="120" />
          <el-table-column align="left" label="createdAt字段" prop="createdAt" width="120" />
         <el-table-column align="left" label="updatedAt字段" prop="updatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.updatedAt) }}</template>
         </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateEarnDailyIncomeMoneyLogFunc(scope.row)">变更</el-button>
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
            <el-form-item label="用户id:"  prop="uid" >
              <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入用户id" />
            </el-form-item>
            <el-form-item label="理财产品ID:"  prop="productId" >
              <el-input v-model.number="formData.productId" :clearable="true" placeholder="请输入理财产品ID" />
            </el-form-item>
            <el-form-item label="下单ID:"  prop="subscribeId" >
              <el-input v-model.number="formData.subscribeId" :clearable="true" placeholder="请输入下单ID" />
            </el-form-item>
            <el-form-item label="当前收益:"  prop="earnings" >
              <el-input-number v-model="formData.earnings"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="当天利率:"  prop="interestRates" >
              <el-input-number v-model="formData.interestRates"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="买入数量:"  prop="boughtNum" >
              <el-input-number v-model="formData.boughtNum"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="模拟|||默认不是模拟:"  prop="isMoni" >
              <el-switch v-model="formData.isMoni" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="到帐时间:"  prop="offeredAt" >
              <el-input v-model.number="formData.offeredAt" :clearable="true" placeholder="请输入到帐时间" />
            </el-form-item>
            <el-form-item label="createdAt字段:"  prop="createdAt" >
              <el-input v-model.number="formData.createdAt" :clearable="true" placeholder="请输入createdAt字段" />
            </el-form-item>
            <el-form-item label="updatedAt字段:"  prop="updatedAt" >
              <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="id字段">
                        {{ detailFrom.id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户id">
                        {{ detailFrom.uid }}
                    </el-descriptions-item>
                    <el-descriptions-item label="理财产品ID">
                        {{ detailFrom.productId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="下单ID">
                        {{ detailFrom.subscribeId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="当前收益">
                        {{ detailFrom.earnings }}
                    </el-descriptions-item>
                    <el-descriptions-item label="当天利率">
                        {{ detailFrom.interestRates }}
                    </el-descriptions-item>
                    <el-descriptions-item label="买入数量">
                        {{ detailFrom.boughtNum }}
                    </el-descriptions-item>
                    <el-descriptions-item label="模拟|||默认不是模拟">
                        {{ detailFrom.isMoni }}
                    </el-descriptions-item>
                    <el-descriptions-item label="到帐时间">
                        {{ detailFrom.offeredAt }}
                    </el-descriptions-item>
                    <el-descriptions-item label="createdAt字段">
                        {{ detailFrom.createdAt }}
                    </el-descriptions-item>
                    <el-descriptions-item label="updatedAt字段">
                        {{ detailFrom.updatedAt }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createEarnDailyIncomeMoneyLog,
  deleteEarnDailyIncomeMoneyLog,
  deleteEarnDailyIncomeMoneyLogByIds,
  updateEarnDailyIncomeMoneyLog,
  findEarnDailyIncomeMoneyLog,
  getEarnDailyIncomeMoneyLogList
} from '@/api/earn/earnDailyIncomeMoneyLog'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'EarnDailyIncomeMoneyLog'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: undefined,
            uid: undefined,
            productId: undefined,
            subscribeId: undefined,
            earnings: 0,
            interestRates: 0,
            boughtNum: 0,
            isMoni: false,
            offeredAt: undefined,
            createdAt: undefined,
            updatedAt: new Date(),
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
    if (searchInfo.value.isMoni === ""){
        searchInfo.value.isMoni=null
    }
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
  const table = await getEarnDailyIncomeMoneyLogList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteEarnDailyIncomeMoneyLogFunc(row)
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
      const res = await deleteEarnDailyIncomeMoneyLogByIds({ ids })
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
const updateEarnDailyIncomeMoneyLogFunc = async(row) => {
    const res = await findEarnDailyIncomeMoneyLog({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteEarnDailyIncomeMoneyLogFunc = async (row) => {
    const res = await deleteEarnDailyIncomeMoneyLog({ id: row.id })
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
        uid: undefined,
        productId: undefined,
        subscribeId: undefined,
        earnings: 0,
        interestRates: 0,
        boughtNum: 0,
        isMoni: false,
        offeredAt: undefined,
        createdAt: undefined,
        updatedAt: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createEarnDailyIncomeMoneyLog(formData.value)
                  break
                case 'update':
                  res = await updateEarnDailyIncomeMoneyLog(formData.value)
                  break
                default:
                  res = await createEarnDailyIncomeMoneyLog(formData.value)
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
  const res = await findEarnDailyIncomeMoneyLog({ id: row.id })
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