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
      client_id: 'CLIENT_ID',
      scope: 'SCOPE',
      ux_mode: 'popup',
      access_type: 'offline',
      callback: (response) => {
        console.log(response);
         const xhr = new XMLHttpRequest();
         xhr.open('POST', 'http://127.0.0.1:8080/token', true);
         xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
         // Set custom header for CRSF
         xhr.setRequestHeader('X-Requested-With', 'XmlHttpRequest');
         xhr.onload = function() {
          console.log('Auth code response: ' + xhr.responseText);
        };
        xhr.send('code=' + response.code);
      },
    });
  },
}
</script>

<template>
  <h2>HomeView</h2>
  <button id="authorize_button" @click="this.handleAuthClick()" class="btn">Authorize</button>
</template>