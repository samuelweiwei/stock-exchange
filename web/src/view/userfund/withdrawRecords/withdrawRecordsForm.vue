
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单ID，唯一标识该提现记录:" prop="orderId">
          <el-input v-model="formData.orderId" :clearable="true"  placeholder="请输入订单ID，唯一标识该提现记录" />
       </el-form-item>
        <el-form-item label="第三方订单ID:" prop="thirdOrderId">
          <el-input v-model.number="formData.thirdOrderId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="会员ID，标识提现会员:" prop="memberId">
          <el-input v-model.number="formData.memberId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="会员手机号:" prop="memberPhone">
          <el-input v-model="formData.memberPhone" :clearable="true"  placeholder="请输入会员手机号" />
       </el-form-item>
        <el-form-item label="提现渠道ID:" prop="withdrawChannelId">
          <el-input v-model.number="formData.withdrawChannelId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="提现货币类型，如BTC:" prop="currency">
          <el-input v-model="formData.currency" :clearable="true"  placeholder="请输入提现货币类型，如BTC" />
       </el-form-item>
        <el-form-item label="提现的货币数量:" prop="withdrawAmount">
          <el-input-number v-model="formData.withdrawAmount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="折算为USDT的金额:" prop="exchangedAmountUsdt">
          <el-input-number v-model="formData.exchangedAmountUsdt" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="提现渠道1 地址 2 银行卡 3 其他:" prop="channelType">
          <el-input v-model="formData.channelType" :clearable="true"  placeholder="请输入提现渠道1 地址 2 银行卡 3 其他" />
       </el-form-item>
        <el-form-item label="提现方式：系统提现，快捷提现:" prop="withdrawType">
          <el-input v-model="formData.withdrawType" :clearable="true"  placeholder="请输入提现方式：系统提现，快捷提现" />
       </el-form-item>
        <el-form-item label="渠道ERC20,TRC20 其他等等:" prop="channel">
          <el-input v-model="formData.channel" :clearable="true"  placeholder="请输入渠道ERC20,TRC20 其他等等" />
       </el-form-item>
        <el-form-item label="用户地址，存储提现发起地址:" prop="userAddress">
          <el-input v-model="formData.userAddress" :clearable="true"  placeholder="请输入用户地址，存储提现发起地址" />
       </el-form-item>
        <el-form-item label="提现目标地址，接收提现的地址:" prop="targetAddress">
          <el-input v-model="formData.targetAddress" :clearable="true"  placeholder="请输入提现目标地址，接收提现的地址" />
       </el-form-item>
        <el-form-item label="订单状态，待审核、已确认、已完成:" prop="orderStatus">
          <el-input v-model="formData.orderStatus" :clearable="true"  placeholder="请输入订单状态，待审核、已确认、已完成" />
       </el-form-item>
        <el-form-item label="用户提现的时间:" prop="withdrawTime">
          <el-date-picker v-model="formData.withdrawTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="审核通过的时间:" prop="approvalTime">
          <el-date-picker v-model="formData.approvalTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="用户操作，提交申请或撤销:" prop="userAction">
          <el-input v-model="formData.userAction" :clearable="true"  placeholder="请输入用户操作，提交申请或撤销" />
       </el-form-item>
        <el-form-item label="审核状态，锁定或解锁:" prop="reviewStatus">
          <el-input v-model="formData.reviewStatus" :clearable="true"  placeholder="请输入审核状态，锁定或解锁" />
       </el-form-item>
        <el-form-item label="1 锁定 0 未锁定:" prop="isLock">
          <el-input v-model="formData.isLock" :clearable="true"  placeholder="请输入1 锁定 0 未锁定" />
       </el-form-item>
        <el-form-item label="拒绝原因:" prop="refusedReason">
          <el-input v-model="formData.refusedReason" :clearable="true"  placeholder="请输入拒绝原因" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createWithdrawRecords,
  updateWithdrawRecords,
  findWithdrawRecords
} from '@/api/userfund/withdrawRecords'

defineOptions({
    name: 'WithdrawRecordsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
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

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWithdrawRecords({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
