<template>
  <div>
    <svg-icon icon-class="question" @click="goto" />
  </div>
</template>

<script>
export default {
  name: 'GoAdminDoc',
  data() {
    return {
      url: 'http://doc.zhangwj.com/kobh-site'
    }
  },
  methods: {
    goto() {
      window.open(this.url)
    }
  }
}
</script>
