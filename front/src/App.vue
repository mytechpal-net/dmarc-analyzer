<script>
import AppNav from './components/AppNav.vue';
export default {
  components: {
    AppNav
  },
}
</script> 
<template>
  <header></header>
  <AppNav />
  <main>
    <div class="flex justify-center">
      <container class="prose">
        <RouterView />
      </container>
    </div>
  </main>
</template>

<style>
@import 'styles/app.css';
</style>
