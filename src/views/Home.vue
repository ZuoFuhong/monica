<template>
  <el-container>
    <el-header>
      <el-row>
        <div class="logo-field">
          <p class="proj-name"><i class="el-icon-s-help"></i><i> MONICA MESH</i></p>
        </div>
      </el-row>
    </el-header>
    <el-container>
      <el-aside>
        <el-row v-for="(item, idx) in menus" :key="idx">
          <p class="menu-item" :class="{'menu-item-active': idx === activeMenu}" @click="switchMenuFocus(idx)"><i :class="item.icon"></i> {{item.name}}</p>
        </el-row>
      </el-aside>
      <el-main>
        <el-row style="text-align: left;">
          <el-form :inline="true" :model="formInline" class="demo-form-inline">
            <el-form-item label="服务名">
              <el-input v-model="formInline.serviceName" placeholder="请输入服务名" maxlength="30"></el-input>
            </el-form-item>
            <el-form-item label="实例IP">
              <el-input v-model="formInline.instance" placeholder="请输入实例IP" maxlength="20"></el-input>
            </el-form-item>
            <el-form-item label="业务">
              <el-input v-model="formInline.businessName" placeholder="请输入业务名" maxlength="20"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button icon="el-icon-search" @click="searchService">查询</el-button>
              <el-button icon="el-icon-plus" @click="addService">新建</el-button>
            </el-form-item>
          </el-form>
        </el-row>
        <el-row>
          <el-table
            :data="serviceList"
            border
            style="width: 100%">
            <el-table-column
              prop="name"
              label="服务名"
              width="180">
            </el-table-column>
            <el-table-column
              prop="namespace"
              label="环境名称"
              width="180">
            </el-table-column>
            <el-table-column
              prop="business"
              label="业务名称">
            </el-table-column>
            <el-table-column
              :formatter="formatTime"
              prop="ctime"
              label="创建时间">
            </el-table-column>
            <el-table-column
              :formatter="formatTime"
              prop="mtime"
              label="修改时间">
            </el-table-column>
            <el-table-column
              fixed="right"
              label="操作"
              width="100">
              <template slot-scope="scope">
                <el-button @click="handleEditService(scope.row)" type="text" size="small">编辑</el-button>
                <el-button @click="handleDeleteService(scope.row)" type="text" size="small">删除</el-button>
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
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import moment from 'moment';
export default {
  name: 'Home',
  data() {
    return {
      // 侧边栏菜单
      activeMenu: 0,
      menus: [
        {
          icon: "el-icon-postcard",
          name: "服务列表"
        },
        {
          icon: "el-icon-coin",
          name: "配置分组"
        }
      ],
      formInline: {
        serviceName: '',
        instance: '',
        businessName: ''
      },
      serviceList: [
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a3',
          name: 'test',
          namespace: 'Test',
          business: '测试应用',
          ctime: 1655716935,
          mtime: 1655716949,
          metadata: {}
        },
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a2',
          name: 'test',
          namespace: 'Test',
          business: '测试应用2',
          ctime: 1655716935,
          mtime: 1655716949,
          metadata: {}
        }
      ],
      total: 100,
      pageSize: 10,
      currentPage: 1
    }
  },
  methods: {
    // 切换菜单
    switchMenuFocus(idx) {
      this.activeMenu = idx
    },
    // 搜索服务
    searchService() {
      console.log('submit!');
    },
    // 新建服务
    addService() {

    },
    // 编辑服务
    handleEditService(row) {
      console.log(row)
    },
    // 删除服务
    handleDeleteService(row) {
      console.log(row)
    },
    // 时间格式
    formatTime(row, column, cellValue) {
      // 转毫秒时间戳
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
.el-container {
  height: 100%;
}
.el-header {
  padding: 0;
  background-color: #262f3e;
  .logo-field {
    width: 199px;
    border-right: 1px solid #000000;
    .proj-name{
      color: #ffffff;
      font-size: 20px;
      font-weight: bold;
      margin: 0;
      line-height: 60px;
    }
  }
}
.el-aside {
  background-color: #1f222c;
  width: 200px !important;
  .menu-item {
    color: #c1c6c8;
    text-align: left;
    padding: 15px 20px;
    margin: 0;
  }
  .menu-item:hover{
    background-color: #2d3546;
    color: #ffffff;
    cursor: pointer;
  }
  .menu-item-active {
    background-color: #2d3546;
    color: #ffffff;
    cursor: pointer;
  }
}
</style>
