
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
            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
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
        
          <el-table-column align="left" label="用户UUID" prop="uuid" width="120" />
          <el-table-column align="left" label="用户登录名" prop="username" width="120" />
          <el-table-column align="left" label="用户登录密码" prop="password" width="120" />
          <el-table-column align="left" label="用户支付密码" prop="paymentPassword" width="120" />
          <el-table-column align="left" label="用户昵称" prop="nickName" width="120" />
          <el-table-column align="left" label="用户头像" prop="headerImg" width="120" />
          <el-table-column align="left" label="用户手机号" prop="phone" width="120" />
          <el-table-column align="left" label="用户邮箱" prop="email" width="120" />
          <el-table-column align="left" label="用户是否被冻结 1正常 2冻结" prop="enable" width="120" />
          <el-table-column align="left" label="直接上级用户ID" prop="parentId" width="120" />
          <el-table-column align="left" label="用户手机号国家码" prop="countryId" width="120" />
        <el-table-column align="left" label="证件类型：1-身份证，2-护照" prop="idType" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.idType) }}</template>
        </el-table-column>
          <el-table-column align="left" label="证件照" prop="idImages" width="120" />
        <el-table-column align="left" label="实名状态：0-未实名，1-待审核， 3-审核通过" prop="authenticationStatus" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.authenticationStatus) }}</template>
        </el-table-column>
          <el-table-column align="left" label="真实姓名" prop="realName" width="120" />
          <el-table-column align="left" label="证件号码" prop="idNumber" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateFrontendUsersFunc(scope.row)">变更</el-button>
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
            <el-form-item label="用户UUID:"  prop="uuid" >
              <el-input v-model="formData.uuid" :clearable="true"  placeholder="请输入用户UUID" />
            </el-form-item>
            <el-form-item label="用户登录名:"  prop="username" >
              <el-input v-model="formData.username" :clearable="true"  placeholder="请输入用户登录名" />
            </el-form-item>
            <el-form-item label="用户登录密码:"  prop="password" >
              <el-input v-model="formData.password" :clearable="true"  placeholder="请输入用户登录密码" />
            </el-form-item>
            <el-form-item label="用户支付密码:"  prop="paymentPassword" >
              <el-input v-model="formData.paymentPassword" :clearable="true"  placeholder="请输入用户支付密码" />
            </el-form-item>
            <el-form-item label="用户昵称:"  prop="nickName" >
              <el-input v-model="formData.nickName" :clearable="true"  placeholder="请输入用户昵称" />
            </el-form-item>
            <el-form-item label="用户头像:"  prop="headerImg" >
              <el-input v-model="formData.headerImg" :clearable="true"  placeholder="请输入用户头像" />
            </el-form-item>
            <el-form-item label="用户手机号:"  prop="phone" >
              <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入用户手机号" />
            </el-form-item>
            <el-form-item label="用户邮箱:"  prop="email" >
              <el-input v-model="formData.email" :clearable="true"  placeholder="请输入用户邮箱" />
            </el-form-item>
            <el-form-item label="用户是否被冻结 1正常 2冻结:"  prop="enable" >
              <el-input v-model.number="formData.enable" :clearable="true" placeholder="请输入用户是否被冻结 1正常 2冻结" />
            </el-form-item>
            <el-form-item label="直接上级用户ID:"  prop="parentId" >
              <el-input v-model.number="formData.parentId" :clearable="true" placeholder="请输入直接上级用户ID" />
            </el-form-item>
            <el-form-item label="用户手机号国家码:"  prop="countryId" >
              <el-input v-model.number="formData.countryId" :clearable="true" placeholder="请输入用户手机号国家码" />
            </el-form-item>
            <el-form-item label="证件类型：1-身份证，2-护照:"  prop="idType" >
              <el-switch v-model="formData.idType" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="证件照:"  prop="idImages" >
              <el-input v-model="formData.idImages" :clearable="true"  placeholder="请输入证件照" />
            </el-form-item>
            <el-form-item label="实名状态：0-未实名，1-待审核， 3-审核通过:"  prop="authenticationStatus" >
              <el-switch v-model="formData.authenticationStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="真实姓名:"  prop="realName" >
              <el-input v-model="formData.realName" :clearable="true"  placeholder="请输入真实姓名" />
            </el-form-item>
            <el-form-item label="证件号码:"  prop="idNumber" >
              <el-input v-model="formData.idNumber" :clearable="true"  placeholder="请输入证件号码" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="用户UUID">
                        {{ detailFrom.uuid }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户登录名">
                        {{ detailFrom.username }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户登录密码">
                        {{ detailFrom.password }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户支付密码">
                        {{ detailFrom.paymentPassword }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户昵称">
                        {{ detailFrom.nickName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户头像">
                        {{ detailFrom.headerImg }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户手机号">
                        {{ detailFrom.phone }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户邮箱">
                        {{ detailFrom.email }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户是否被冻结 1正常 2冻结">
                        {{ detailFrom.enable }}
                    </el-descriptions-item>
                    <el-descriptions-item label="直接上级用户ID">
                        {{ detailFrom.parentId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户手机号国家码">
                        {{ detailFrom.countryId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="证件类型：1-身份证，2-护照">
                        {{ detailFrom.idType }}
                    </el-descriptions-item>
                    <el-descriptions-item label="证件照">
                        {{ detailFrom.idImages }}
                    </el-descriptions-item>
                    <el-descriptions-item label="实名状态：0-未实名，1-待审核， 3-审核通过">
                        {{ detailFrom.authenticationStatus }}
                    </el-descriptions-item>
                    <el-descriptions-item label="真实姓名">
                        {{ detailFrom.realName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="证件号码">
                        {{ detailFrom.idNumber }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createFrontendUsers,
  deleteFrontendUsers,
  deleteFrontendUsersByIds,
  updateFrontendUsers,
  findFrontendUsers,
  getFrontendUsersList
} from '@/api/user/frontendUsers'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'FrontendUsers'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            uuid: '',
            username: '',
            password: '',
            paymentPassword: '',
            nickName: '',
            headerImg: '',
            phone: '',
            email: '',
            enable: undefined,
            parentId: undefined,
            countryId: undefined,
            idType: false,
            idImages: '',
            authenticationStatus: false,
            realName: '',
            idNumber: '',
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
    if (searchInfo.value.idType === ""){
        searchInfo.value.idType=null
    }
    if (searchInfo.value.authenticationStatus === ""){
        searchInfo.value.authenticationStatus=null
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
  const table = await getFrontendUsersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteFrontendUsersFunc(row)
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
      const res = await deleteFrontendUsersByIds({ IDs })
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
const updateFrontendUsersFunc = async(row) => {
    const res = await findFrontendUsers({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteFrontendUsersFunc = async (row) => {
    const res = await deleteFrontendUsers({ ID: row.ID })
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
        uuid: '',
        username: '',
        password: '',
        paymentPassword: '',
        nickName: '',
        headerImg: '',
        phone: '',
        email: '',
        enable: undefined,
        parentId: undefined,
        countryId: undefined,
        idType: false,
        idImages: '',
        authenticationStatus: false,
        realName: '',
        idNumber: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createFrontendUsers(formData.value)
                  break
                case 'update':
                  res = await updateFrontendUsers(formData.value)
                  break
                default:
                  res = await createFrontendUsers(formData.value)
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
  const res = await findFrontendUsers({ ID: row.ID })
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