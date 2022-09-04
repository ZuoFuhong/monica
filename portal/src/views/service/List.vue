<template>
  <div>
    <el-row class="content-header">
      <p class="page-title">服务列表</p>
    </el-row>
    <div class="content-body">
      <el-row>
        <el-form :inline="true" :model="formInline">
          <el-form-item label="服务名">
            <el-input v-model="formInline.serviceName" placeholder="请输入服务名" maxlength="30"></el-input>
          </el-form-item>
          <el-form-item label="业务">
            <el-input v-model="formInline.businessName" placeholder="请输入业务名" maxlength="20"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button icon="el-icon-search" @click="handleSearch">查询</el-button>
            <el-button icon="el-icon-plus" @click="handleAdd">新建</el-button>
          </el-form-item>
        </el-form>
      </el-row>
      <el-row>
        <el-table
          :data="serviceList"
          border
          style="width: 100%">
          <el-table-column
            show-overflow-tooltip
            prop="name"
            label="服务名">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="business"
            label="业务名称">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="owners"
            label="负责人">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="remark"
            label="备注">
          </el-table-column>
          <el-table-column
            :formatter="formatTime"
            show-overflow-tooltip
            prop="ctime"
            label="创建时间">
          </el-table-column>
          <el-table-column
            :formatter="formatTime"
            show-overflow-tooltip
            prop="mtime"
            label="修改时间">
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
              <el-button @click="handleEdit(scope.row)" type="text" size="small">编辑</el-button>
              <el-button @click="handleDelete(scope.row)" type="text" size="small"><span class="deleteBtn">删除</span></el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-row>
      <el-row style="text-align: right; padding: 20px 0;">
        <el-pagination
          @current-change="handleCurrentChange"
          :current-page.sync="currentPage"
          :page-size="pageSize"
          layout="total, prev, pager, next"
          :total="total">
        </el-pagination>
      </el-row>
    </div>
    <!-- 编辑框 -->
    <el-dialog
      :title="dialogTitle"
      :visible.sync="dialogVisible"
      width="40%">
      <el-row>
        <el-col :span="22">
          <el-form :model="dialogForm" ref="dialogForm" label-width="100px">
            <el-form-item label="服务名称" prop="name" class="is-required">
              <el-input v-model="dialogForm.name"></el-input>
            </el-form-item>
            <el-form-item label="业务名称" prop="business" class="is-required">
              <el-input v-model="dialogForm.business"></el-input>
            </el-form-item>
            <el-form-item label="负责人" prop="owners" class="is-required">
              <el-input v-model="dialogForm.owners"></el-input>
            </el-form-item>
            <el-form-item label="备注" prop="remark">
              <el-input v-model="dialogForm.remark"></el-input>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="doSaveService">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import moment from 'moment';
import serv from '../../model/service'
export default {
  name: "Service",
  created() {
    this.refreshTableList()
  },
  data() {
    return {
      formInline: {
        serviceName: '',
        businessName: ''
      },
      serviceList: [],
      total: 100,
      pageSize: 10,
      currentPage: 1,
      // 编辑框
      dialogVisible: false,
      dialogTitle: '',
      dialogForm: {
        id: 0,
        name: '',
        business: '',
        owners: '',
        remark: ''
      }
    }
  },
  methods: {
    // 新建服务
    handleAdd() {
      this.dialogVisible = true
      this.dialogTitle = "新建服务"
      this.dialogForm = {
        id: 0,
        name: '',
        business: '',
        owners: '',
        remark: ''
      }
    },
    // 编辑服务
    handleEdit(row) {
      this.dialogVisible = true
      this.dialogTitle = "编辑服务"
      this.dialogForm = {
        id: row.id,
        name: row.name,
        business: row.business,
        owners: row.owners,
        remark: row.remark
      }
    },
    // 保存提交
    async doSaveService() {
      let form = this.dialogForm
      let {retcode, errmsg} = await serv.saveService({
        id: form.id,
        name: form.name,
        business: form.business,
        owners: form.owners,
        remark: form.remark
      })
      if (retcode === 0) {
        this.$message({
          message: '保存成功',
          type: 'success'
        });
        this.dialogVisible = false
        this.refreshTableList()
        return
      }
      this.$message.error(errmsg)
    },
    // 删除记录
    async handleDelete(row) {
      let {retcode, errmsg} = await serv.deleteService(row.name)
      if (retcode === 0) {
        this.$message({
          type: 'success',
          message: '删除成功!'
        });
        this.refreshTableList()
        return
      }
      this.$message.error(errmsg)
    },
    // 搜索
    async handleSearch() {
      let {retcode, errmsg, data} = await serv.searchService({
        kw: this.formInline.serviceName,
        busi: this.formInline.businessName,
        page: this.currentPage,
        page_size: this.pageSize
      })
      if (retcode !== 0) {
        this.$message.error(errmsg)
        return
      }
      this.serviceList = data.list
      this.total = data.total
    },
    // 切换分页
    async handleCurrentChange(val) {
      this.currentPage = val
      this.handleSearch()
    },
    // 刷新列表
    refreshTableList() {
      this.currentPage = 1
      this.handleSearch()
    },
    // 时间格式化
    formatTime(row, column, cellValue) {
      if (cellValue === 0) {
        return "-";
      }
      return moment(cellValue * 1000).format("YYYY-MM-DD HH:mm:ss");
    }
  }
}
</script>

<style lang="scss" scoped>
.deleteBtn{
  color: red;
}
.deleteBtn:hover{
  color: rgb(251, 100, 100);
}
</style>