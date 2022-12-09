<template>
  <div>
    <el-row class="content-header">
      <p class="page-title">服务实例 {{namespace}}/{{service_name}}</p>
    </el-row>
    <div class="content-body">
      <el-row>
        <el-form :inline="true" :model="formInline">
          <el-form-item label="实例IP">
            <el-input v-model="formInline.ip" placeholder="请输入实例IP" maxlength="20"></el-input>
          </el-form-item>
          <el-form-item label="端口">
            <el-input v-model="formInline.port" placeholder="请输入端口" maxlength="5"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button icon="el-icon-search" @click="handleSearch">查询</el-button>
            <el-button icon="el-icon-plus" @click="handleAdd">新增服务实例</el-button>
          </el-form-item>
        </el-form>
      </el-row>
      <el-row>
        <el-table
          :data="instanceList"
          border
          style="width: 100%">
          <el-table-column
            show-overflow-tooltip
            prop="ip"
            label="实例IP">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="port"
            label="端口">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="weight"
            label="权重">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="healthy"
            label="健康状态">
            <template slot-scope="scope">
              <span style="color: red;" v-if="scope.row.healthy === 1">异常</span>
              <span style="color: #3ecc36;" v-if="scope.row.healthy === 2">健康</span>
            </template>
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="isolate"
            label="隔离状态">
            <template slot-scope="scope">
              <span style="color: red;" v-if="scope.row.isolate === 2">隔离</span>
              <span style="color: #3ecc36;" v-if="scope.row.isolate === 1">不隔离</span>
            </template>
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
            <el-form-item label="实例IP" prop="ip" class="is-required">
              <el-input v-model="dialogForm.ip" maxlength="20" :disabled="dialogForm.action === 1"></el-input>
            </el-form-item>
            <el-form-item label="端口" prop="port" class="is-required">
              <el-input v-model="dialogForm.port" maxlength="5" :disabled="dialogForm.action === 1"></el-input>
            </el-form-item>
            <el-form-item label="权重" prop="weight" class="is-required">
              <el-input v-model="dialogForm.weight" maxlength="3"></el-input>
            </el-form-item>
            <el-form-item label="是否隔离" prop="isolate" class="is-required">
              <el-switch
                v-model="dialogForm.isolate"
                active-color="#ff4949"
                inactive-color="#13ce66">
              </el-switch>
            </el-form-item>
            <el-form-item label="元数据" prop="remark">
              <el-input v-model="dialogForm.metadata" type="textarea" rows="3" placeholder="元数据是服务实例的附加信息，可用于标记实例特征" maxlength="100" show-word-limit></el-input>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="doSaveInstance">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import moment from 'moment';
import serv from '../../model/service'
export default {
  name: "Instance",
  created() {
    this.namespace = this.$route.params.ns
    this.service_name = this.$route.params.name
    this.handleSearch()
  },
  data() {
    return {
      namespace: "",
      service_name: "",
      formInline: {
        ip: "",
        port: ""
      },
      // 服务实例
      instanceList: [],
      total: 100,
      pageSize: 10,
      currentPage: 1,
      // 编辑框
      dialogVisible: false,
      dialogTitle: '',
      dialogForm: {
        id: 0,
        ns: "",
        sname: "",
        isolate: false,
        ip: "",
        port: '',
        weight: '',
        metadata: '',
        // 动作: 0-新增 1-修改
        action: 0
      }
    }
  },
  methods: {
    // 搜索服务实例
    async handleSearch() {
      let port = 0
      if (this.formInline.port) {
        port = parseInt(this.formInline.port)
      }
      let {retcode, errmsg, data} = await serv.searchServiceInstance({
        ns: this.namespace,
        sname: this.service_name,
        ip: this.formInline.ip,
        port: port,
        page: this.currentPage,
        page_size: this.pageSize
      })
      if (retcode !== 0) {
        this.$message.error(errmsg)
        return
      }
      this.instanceList = data.list
      this.total = data.total
    },
    // 新建实例
    handleAdd() {
      this.dialogVisible = true
      this.dialogTitle = '新增服务实例'
      this.dialogForm = {
        id: 0,
        ns: this.namespace,
        sname: this.service_name
      }
    },
    // 编辑实例
    handleEdit(row) {
      this.dialogVisible = true
      this.dialogTitle = '编辑服务实例'
      this.dialogForm = {
        id: row.id,
        ns: row.namespace,
        sname: row.service_name,
        isolate: row.isolate === 2,
        ip: row.ip,
        port: row.port,
        weight: row.weight,
        metadata: row.metadata,
        action: 1
      }
    },
    // 删除实例
    handleDelete(row) {
      let that = this;
      this.$confirm(`此操作将删除服务实例 ${row.ip}:${row.port}, 是否继续?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        // 删除请求
        let {retcode, errmsg} = await serv.deleteServiceInstance(row.namespace, row.service_name, row.ip)
        if (retcode !== 0) {
          this.$message.error(errmsg)
          return
        }
        this.$message({
          type: 'success',
          message: '删除成功!'
        });
        // 刷新列表
        that.handleSearch()
      }).catch(() => {
        // cancel delete, nothing to do         
      });
    },
    // 保存示例
    async doSaveInstance() {
      let {retcode, errmsg} = await serv.saveServiceInstance({
        id: this.dialogForm.id,
        ns: this.dialogForm.ns,
        sname: this.dialogForm.sname,
        isolate: this.dialogForm.isolate ? 2 : 1,
        ip: this.dialogForm.ip,
        port: parseInt(this.dialogForm.port),
        weight: parseInt(this.dialogForm.weight),
        metadata: this.dialogForm.metadata
      })
      if (retcode !== 0) {
        this.$message.error(errmsg)
        return
      }
      this.$message({
        type: 'success',
        message: '提交成功!'
      });
      // 关闭
      this.dialogVisible = false
      // 刷新
      this.handleSearch()
    },
    // 切换分页
    handleCurrentChange(val) {
      this.currentPage = val
      this.handleSearch()
    },
    // 时间格式化
    formatTime(row, column, cellValue) {
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