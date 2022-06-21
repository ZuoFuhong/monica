<template>
  <div>
    <el-row class="content-header">
      <p class="page-title">服务实例 {{namespace}}/{{service_name}}</p>
    </el-row>
    <div class="content-body">
      <el-row>
        <el-form :inline="true" :model="formInline">
          <el-form-item label="实例IP">
            <el-input v-model="formInline.host" placeholder="请输入实例IP" maxlength="20"></el-input>
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
            prop="host"
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
              <span style="color: red;" v-if="!scope.row.healthy">异常</span>
              <span style="color: #3ecc36;" v-if="scope.row.healthy">健康</span>
            </template>
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="isolate"
            label="隔离状态">
            <template slot-scope="scope">
              <span style="color: red;" v-if="scope.row.isolate">隔离</span>
              <span style="color: #3ecc36;" v-if="!scope.row.isolate">不隔离</span>
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
  </div>
</template>

<script>
import moment from 'moment';

export default {
  name: "Instance",
  created() {
    this.namespace = this.$route.params.ns
    this.service_name = this.$route.params.name
  },
  data() {
    return {
      namespace: "",
      service_name: "",
      formInline: {
        host: "",
        port: ""
      },
      // 服务实例
      instanceList: [
        {
          id: "7964f48ff4c24c39e4b006cefbfb4b63f0236219",
          healthy: true,
          isolate: false,
          host: "192.168.1.1",
          metadata: {},
          port: 8080,
          service: "test",
          weight: 100,
          ctime: 1655716935,
          mtime: 1655716949,
          namespace: 'Test'
        },
        {
          id: "7864f48ff4c24c39e4b006cefbfb4b63f0236219",
          healthy: false,
          isolate: true,
          host: "192.168.1.2",
          metadata: {},
          port: 8080,
          service: "test2",
          weight: 100,
          ctime: 1655716935,
          mtime: 1655716949,
          namespace: 'Test'
        }
      ],
      total: 100,
      pageSize: 10,
      currentPage: 1
    }
  },
  methods: {
    // 搜索实例
    handleSearch() {
      console.log('submit!');
    },
    // 新建实例
    handleAdd() {
    },
    // 编辑服务
    handleEdit(row) {
      console.log(row)
    },
    // 删除服务
    handleDelete(row) {
      console.log(row)
    },
    // 切换分页
    handleCurrentChange(val) {
      // Todo: 分页加载
      console.log(val)
    },
    // 时间格式
    formatTime(row, column, cellValue) {
      // 转毫秒时间戳
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