<script>
export default {
  setup() {
  },
  data() {
    return {
      googleClient: null,
    }
  },
  methods: {
    handleAuthClick() {
      console.log('authorize')
      this.googleClient.requestCode()
    },
    handleSignoutClick() {
      console.log('sign out')
    },
  },
  mounted() {
    this.googleClient = google.accounts.oauth2.initCodeClient({
      client_id: 'INSERT_CLIENT_ID',
      scope: 'REPLACE_WITH_SCOPE',
      ux_mode: 'popup',
      access_type: 'offline',
      callback: (response) => {
        console.log(response);
      },
    });
  },
}
</script>

<template>
  <h2>HomeView</h2>
  <button id="authorize_button" @click="this.handleAuthClick()" class="btn">Authorize</button>
</template>