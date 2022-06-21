<template>
  <el-container>
    <el-header>
      <el-row>
        <div class="logo-field">
          <p class="proj-name"><i class="el-icon-s-help"></i><i @click="backToHomePage"> MONICA MESH</i></p>
        </div>
      </el-row>
    </el-header>
    <el-container>
      <el-aside>
        <el-row v-for="(item, idx) in menus" :key="idx">
          <p class="menu-item" :class="{'menu-item-active': idx === activeMenu}" @click="switchMenuFocus(item, idx)">
            <i :class="item.icon"></i> {{item.name}}
          </p>
        </el-row>
      </el-aside>
      <el-main>
        <router-view/>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      // 菜单导航
      activeMenu: 0,
      menus: [
        {
          path: "/service",
          icon: "el-icon-menu",
          name: "服务列表"
        },
        {
          path: "/naming-service",
          icon: "el-icon-postcard",
          name: "服务治理"
        },
        {
          path: "/kvgroup",
          icon: "el-icon-coin",
          name: "配置中心"
        }
      ]
    }
  },
  computed: {
    activeMenuTitle() {
      return this.menus[this.activeMenu].name
    }
  },
  methods: {
    // 菜单切换
    switchMenuFocus(item, idx) {
      this.activeMenu = idx
      this.$router.push(item.path)
    },
    backToHomePage() {
      this.$router.push('/')
    }
  }
}
</script>

<style lang="scss">
html,
body {
  margin: 0;
  padding: 0;
  height: 100%;
}
.el-container {
  height: 100%;
}
.el-header {
  padding: 0 !important;
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
      text-align: center;
    }
    .proj-name:hover {
      cursor: pointer;
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
.el-main {
  padding: 0 !important;
  background-color: #f3f4f7;
  .content-header {
    background-color: #ffffff;
    .page-title {
      padding-left: 20px;
      text-align: left;
      font-weight: bold;
    }
  }
  .content-body {
    text-align: left;
    padding: 20px;
  }
}
</style>
