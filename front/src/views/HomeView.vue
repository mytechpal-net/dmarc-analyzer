<script>
export default {
  setup() {
  },
  methods: {
    goToAbout() {
      this.$router.push('/about')
    },
  },
  mounted() {
    if (window.gapi) {
      console.log(window.gapi)
    }
  },
}
</script>

<template>
  <h2>HomeView</h2>
  <button @click="goToAbout">Go to About</button>
</template>