
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户id:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="理财产品ID:" prop="productId">
          <el-input v-model.number="formData.productId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="违约金比例:" prop="penaltyRatio">
          <el-input-number v-model="formData.penaltyRatio" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="模拟|||默认不是模拟:" prop="isMoni">
          <el-switch v-model="formData.isMoni" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="状态|||1: 质押,2:赎回:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="1提前赎回, 2未提前赎回:" prop="redeemInAdvance">
          <el-switch v-model="formData.redeemInAdvance" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="罚金:" prop="fine">
          <el-input-number v-model="formData.fine" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="活/定期产品开始时间:" prop="startAt">
          <el-input v-model.number="formData.startAt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="活/定期产品到期时间:" prop="endAt">
          <el-input v-model.number="formData.endAt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="买入数量:" prop="boughtNum">
          <el-input-number v-model="formData.boughtNum" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="createdAt字段:" prop="createdAt">
          <el-date-picker v-model="formData.createdAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="updatedAt字段:" prop="updatedAt">
          <el-date-picker v-model="formData.updatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createEarnSubscribeLog,
  updateEarnSubscribeLog,
  findEarnSubscribeLog
} from '@/api/earn/earnSubscribeLog'

defineOptions({
    name: 'EarnSubscribeLogForm'
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
            id: undefined,
            uid: undefined,
            productId: undefined,
            penaltyRatio: 0,
            isMoni: false,
            status: false,
            redeemInAdvance: false,
            fine: 0,
            startAt: undefined,
            endAt: undefined,
            boughtNum: 0,
            createdAt: new Date(),
            updatedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEarnSubscribeLog({ ID: route.query.id })
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
               res = await createEarnSubscribeLog(formData.value)
               break
             case 'update':
               res = await updateEarnSubscribeLog(formData.value)
               break
             default:
               res = await createEarnSubscribeLog(formData.value)
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
