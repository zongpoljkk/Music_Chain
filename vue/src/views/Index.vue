<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <!-- <sp-token-send /> -->
      <!-- this line is used by starport scaffolding # 4 -->

    <!-- Upload File -->
    <input type="file" id="file" ref="file" style="display:none" v-on:change="handleFileUpload()"/> 
    <div id="createMusic">
      <button type="button" @click="uploadMusic">Upload File</button>
    </div>
		<sp-type-form type="musics" :fields="['mediaLink', 'price', 'name', ]" module="MusicChain" />
		<sp-type-form type="artist" :fields="['name', ]" module="MusicChain" />
    </div>
  </div>
</template>

<script>

import * as sp from "@tendermint/vue";
import Vue from "vue";
import axios from 'axios';

// window.onload = function () {
// window.addEventListener('load', function () {
// var createMusic = new Vue({
//   el: '#createMusic',
//   methods: {
//     uploadMusic: function () {
//       console.log("yo");
//       // axios.get("http://localhost:12345/MusicChain/musics/1")
//       axios.get('http://localhost:12345/MusicChain/musics').then((res) => {
//           console.log(res);
//         })
//     }
//   }
// })
// }
// }


export default {
  components: { ...sp },
   data() {
     return {
       file: ''
    }
  },
  methods: {
    uploadMusic: function uploadMusic() {
      console.log("yo");
      document.getElementById("file").click()
      // document.getElementById("imgupload").onchange = 
      // let formData = new FormData();
      // formData.append('file', this.file);
      /*
        Make the request to the POST /multiple-files URL
      */
      // axios.post( '/multiple-files',
      //   formData,
      //   {
      //     headers: {
      //         'Content-Type': 'multipart/form-data'
      //     }
      //   }
      // ).then(function(){
      //   console.log('SUCCESS!!');
      // })
      // .catch(function(){
      //   console.log('FAILURE!!');
      // });

      this.handleFileUpload()

      // axios.get("http://localhost:12345/MusicChain/musics/1")
      axios.get('http://localhost:1317/MusicChain/musics').then((res) => {
          console.log(res);
        })
    },
    handleFileUpload: function handleFileUpload() {
      // console.log(this.$refs.file.files[0])
      this.file = this.$refs.file.files[0];
      console.log(this.file)
    },
  },
 
};
</script>
