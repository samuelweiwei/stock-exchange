
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
        
          <el-table-column align="left" label="订单ID，唯一标识该提现记录" prop="orderId" width="120" />
          <el-table-column align="left" label="第三方订单ID" prop="thirdOrderId" width="120" />
          <el-table-column align="left" label="会员ID，标识提现会员" prop="memberId" width="120" />
          <el-table-column align="left" label="会员手机号" prop="memberPhone" width="120" />
          <el-table-column align="left" label="提现渠道ID" prop="withdrawChannelId" width="120" />
          <el-table-column align="left" label="提现货币类型，如BTC" prop="currency" width="120" />
          <el-table-column align="left" label="提现的货币数量" prop="withdrawAmount" width="120" />
          <el-table-column align="left" label="折算为USDT的金额" prop="exchangedAmountUsdt" width="120" />
          <el-table-column align="left" label="提现渠道1 地址 2 银行卡 3 其他" prop="channelType" width="120" />
          <el-table-column align="left" label="提现方式：系统提现，快捷提现" prop="withdrawType" width="120" />
          <el-table-column align="left" label="渠道ERC20,TRC20 其他等等" prop="channel" width="120" />
          <el-table-column align="left" label="用户地址，存储提现发起地址" prop="userAddress" width="120" />
          <el-table-column align="left" label="提现目标地址，接收提现的地址" prop="targetAddress" width="120" />
          <el-table-column align="left" label="订单状态，待审核、已确认、已完成" prop="orderStatus" width="120" />
         <el-table-column align="left" label="用户提现的时间" prop="withdrawTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.withdrawTime) }}</template>
         </el-table-column>
         <el-table-column align="left" label="审核通过的时间" prop="approvalTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.approvalTime) }}</template>
         </el-table-column>
          <el-table-column align="left" label="用户操作，提交申请或撤销" prop="userAction" width="120" />
          <el-table-column align="left" label="审核状态，锁定或解锁" prop="reviewStatus" width="120" />
          <el-table-column align="left" label="1 锁定 0 未锁定" prop="isLock" width="120" />
          <el-table-column align="left" label="拒绝原因" prop="refusedReason" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateWithdrawRecordsFunc(scope.row)">变更</el-button>
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
            <el-form-item label="订单ID，唯一标识该提现记录:"  prop="orderId" >
              <el-input v-model="formData.orderId" :clearable="true"  placeholder="请输入订单ID，唯一标识该提现记录" />
            </el-form-item>
            <el-form-item label="第三方订单ID:"  prop="thirdOrderId" >
              <el-input v-model.number="formData.thirdOrderId" :clearable="true" placeholder="请输入第三方订单ID" />
            </el-form-item>
            <el-form-item label="会员ID，标识提现会员:"  prop="memberId" >
              <el-input v-model.number="formData.memberId" :clearable="true" placeholder="请输入会员ID，标识提现会员" />
            </el-form-item>
            <el-form-item label="会员手机号:"  prop="memberPhone" >
              <el-input v-model="formData.memberPhone" :clearable="true"  placeholder="请输入会员手机号" />
            </el-form-item>
            <el-form-item label="提现渠道ID:"  prop="withdrawChannelId" >
              <el-input v-model.number="formData.withdrawChannelId" :clearable="true" placeholder="请输入提现渠道ID" />
            </el-form-item>
            <el-form-item label="提现货币类型，如BTC:"  prop="currency" >
              <el-input v-model="formData.currency" :clearable="true"  placeholder="请输入提现货币类型，如BTC" />
            </el-form-item>
            <el-form-item label="提现的货币数量:"  prop="withdrawAmount" >
              <el-input-number v-model="formData.withdrawAmount"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="折算为USDT的金额:"  prop="exchangedAmountUsdt" >
              <el-input-number v-model="formData.exchangedAmountUsdt"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="提现渠道1 地址 2 银行卡 3 其他:"  prop="channelType" >
              <el-input v-model="formData.channelType" :clearable="true"  placeholder="请输入提现渠道1 地址 2 银行卡 3 其他" />
            </el-form-item>
            <el-form-item label="提现方式：系统提现，快捷提现:"  prop="withdrawType" >
              <el-input v-model="formData.withdrawType" :clearable="true"  placeholder="请输入提现方式：系统提现，快捷提现" />
            </el-form-item>
            <el-form-item label="渠道ERC20,TRC20 其他等等:"  prop="channel" >
              <el-input v-model="formData.channel" :clearable="true"  placeholder="请输入渠道ERC20,TRC20 其他等等" />
            </el-form-item>
            <el-form-item label="用户地址，存储提现发起地址:"  prop="userAddress" >
              <el-input v-model="formData.userAddress" :clearable="true"  placeholder="请输入用户地址，存储提现发起地址" />
            </el-form-item>
            <el-form-item label="提现目标地址，接收提现的地址:"  prop="targetAddress" >
              <el-input v-model="formData.targetAddress" :clearable="true"  placeholder="请输入提现目标地址，接收提现的地址" />
            </el-form-item>
            <el-form-item label="订单状态，待审核、已确认、已完成:"  prop="orderStatus" >
              <el-input v-model="formData.orderStatus" :clearable="true"  placeholder="请输入订单状态，待审核、已确认、已完成" />
            </el-form-item>
            <el-form-item label="用户提现的时间:"  prop="withdrawTime" >
              <el-date-picker v-model="formData.withdrawTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="审核通过的时间:"  prop="approvalTime" >
              <el-date-picker v-model="formData.approvalTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="用户操作，提交申请或撤销:"  prop="userAction" >
              <el-input v-model="formData.userAction" :clearable="true"  placeholder="请输入用户操作，提交申请或撤销" />
            </el-form-item>
            <el-form-item label="审核状态，锁定或解锁:"  prop="reviewStatus" >
              <el-input v-model="formData.reviewStatus" :clearable="true"  placeholder="请输入审核状态，锁定或解锁" />
            </el-form-item>
            <el-form-item label="1 锁定 0 未锁定:"  prop="isLock" >
              <el-input v-model="formData.isLock" :clearable="true"  placeholder="请输入1 锁定 0 未锁定" />
            </el-form-item>
            <el-form-item label="拒绝原因:"  prop="refusedReason" >
              <el-input v-model="formData.refusedReason" :clearable="true"  placeholder="请输入拒绝原因" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="订单ID，唯一标识该提现记录">
                        {{ detailFrom.orderId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="第三方订单ID">
                        {{ detailFrom.thirdOrderId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="会员ID，标识提现会员">
                        {{ detailFrom.memberId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="会员手机号">
                        {{ detailFrom.memberPhone }}
                    </el-descriptions-item>
                    <el-descriptions-item label="提现渠道ID">
                        {{ detailFrom.withdrawChannelId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="提现货币类型，如BTC">
                        {{ detailFrom.currency }}
                    </el-descriptions-item>
                    <el-descriptions-item label="提现的货币数量">
                        {{ detailFrom.withdrawAmount }}
                    </el-descriptions-item>
                    <el-descriptions-item label="折算为USDT的金额">
                        {{ detailFrom.exchangedAmountUsdt }}
                    </el-descriptions-item>
                    <el-descriptions-item label="提现渠道1 地址 2 银行卡 3 其他">
                        {{ detailFrom.channelType }}
                    </el-descriptions-item>
                    <el-descriptions-item label="提现方式：系统提现，快捷提现">
                        {{ detailFrom.withdrawType }}
                    </el-descriptions-item>
                    <el-descriptions-item label="渠道ERC20,TRC20 其他等等">
                        {{ detailFrom.channel }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户地址，存储提现发起地址">
                        {{ detailFrom.userAddress }}
                    </el-descriptions-item>
                    <el-descriptions-item label="提现目标地址，接收提现的地址">
                        {{ detailFrom.targetAddress }}
                    </el-descriptions-item>
                    <el-descriptions-item label="订单状态，待审核、已确认、已完成">
                        {{ detailFrom.orderStatus }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户提现的时间">
                        {{ detailFrom.withdrawTime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="审核通过的时间">
                        {{ detailFrom.approvalTime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户操作，提交申请或撤销">
                        {{ detailFrom.userAction }}
                    </el-descriptions-item>
                    <el-descriptions-item label="审核状态，锁定或解锁">
                        {{ detailFrom.reviewStatus }}
                    </el-descriptions-item>
                    <el-descriptions-item label="1 锁定 0 未锁定">
                        {{ detailFrom.isLock }}
                    </el-descriptions-item>
                    <el-descriptions-item label="拒绝原因">
                        {{ detailFrom.refusedReason }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createWithdrawRecords,
  deleteWithdrawRecords,
  deleteWithdrawRecordsByIds,
  updateWithdrawRecords,
  findWithdrawRecords,
  getWithdrawRecordsList
} from '@/api/userfund/withdrawRecords'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'




defineOptions({
    name: 'WithdrawRecords'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            orderId: '',
            thirdOrderId: undefined,
            memberId: undefined,
            memberPhone: '',
            withdrawChannelId: undefined,
            currency: '',
            withdrawAmount: 0,
            exchangedAmountUsdt: 0,
            channelType: '',
            withdrawType: '',
            channel: '',
            userAddress: '',
            targetAddress: '',
            orderStatus: '',
            withdrawTime: new Date(),
            approvalTime: new Date(),
            userAction: '',
            reviewStatus: '',
            isLock: '',
            refusedReason: '',
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
  const table = await getWithdrawRecordsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteWithdrawRecordsFunc(row)
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
      const res = await deleteWithdrawRecordsByIds({ IDs })
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
const updateWithdrawRecordsFunc = async(row) => {
    const res = await findWithdrawRecords({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWithdrawRecordsFunc = async (row) => {
    const res = await deleteWithdrawRecords({ ID: row.ID })
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
        orderId: '',
        thirdOrderId: undefined,
        memberId: undefined,
        memberPhone: '',
        withdrawChannelId: undefined,
        currency: '',
        withdrawAmount: 0,
        exchangedAmountUsdt: 0,
        channelType: '',
        withdrawType: '',
        channel: '',
        userAddress: '',
        targetAddress: '',
        orderStatus: '',
        withdrawTime: new Date(),
        approvalTime: new Date(),
        userAction: '',
        reviewStatus: '',
        isLock: '',
        refusedReason: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createWithdrawRecords(formData.value)
                  break
                case 'update':
                  res = await updateWithdrawRecords(formData.value)
                  break
                default:
                  res = await createWithdrawRecords(formData.value)
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
  const res = await findWithdrawRecords({ ID: row.ID })
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