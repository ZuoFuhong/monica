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
export default {
  name: "AppList",
  created() {
    console.log("init service page")
  },
  data() {
    return {
      formInline: {
        serviceName: '',
        businessName: ''
      },
      serviceList: [
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a3',
          name: 'test',
          business: '测试应用',
          owners: 'mars',
          remark: "123",
          ctime: 1655716935,
          mtime: 1655716949
        },
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a2',
          name: 'test',
          business: '测试应用2',
          owners: 'mars',
          remark: "123",
          ctime: 1655716935,
          mtime: 1655716949
        }
      ],
      total: 100,
      pageSize: 10,
      currentPage: 1,
      // 编辑框
      dialogVisible: false,
      dialogTitle: '',
      dialogForm: {
        name: '',
        business: '',
        owners: '',
        remark: ''
      }
    }
  },
  methods: {
    // 搜索服务
    handleSearch() {
      console.log('submit!');
    },
    // 新建服务
    handleAdd() {
      this.dialogVisible = true
      this.dialogTitle = "新建服务"
      this.dialogForm = {}
    },
    // 编辑服务
    handleEdit(row) {
      this.dialogVisible = true
      this.dialogTitle = "编辑服务"
      this.dialogForm = {
        name: row.name,
        business: row.business,
        owners: row.owners,
        remark: row.remark
      }
    },
    // 保存提交
    doSaveService() {
      // Todo: 保存提交，刷新列表
      console.log(this.dialogForm)
    },
    // 删除服务
    handleDelete(row) {
      this.$confirm(`此操作将删除 ${row.name} 服务, 是否继续?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // Todo: 检查使用状态，是否允许删除
        this.$message({
          type: 'success',
          message: '删除成功!'
        });
      }).catch(() => {
        // cancel delete, nothing to do         
      });
    },
    // 时间格式化
    formatTime(row, column, cellValue) {
      return moment(cellValue * 1000).format("YYYY-MM-DD HH:mm:ss");
    },
    // 切换分页
    handleCurrentChange(val) {
      // Todo: 分页加载
      console.log(val)
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