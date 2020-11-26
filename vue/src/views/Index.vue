<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <!-- <sp-token-send /> -->
      <!-- this line is used by starport scaffolding # 4 -->

      <!-- <sp-type-form
        type="musics"
        :fields="['price', 'name', 'file']"
        module="MusicChain"
      /> -->
      <sp-type-form type="artist" :fields="['name']" module="MusicChain" />
      <sp-type-form type="musics" module="MusicChain" />
      <h3>New Music</h3>
      <div style="border:1px solid black;padding:20px">
        <input placeholder="price" id="price" autocomplete="off" /><br />
        <input placeholder="name" id="name" autocomplete="off" /><br />
        <input type="file" id="file" /><br />
        <button v-on:click="onClickCreate()">CREATE MUSIC</button>
      </div>
      <sp-type-form
        type="purchased"
        :fields="['MusicID']"
        module="MusicChain"
      />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";
import axios from "axios";
import {
  Secp256k1Wallet,
  SigningCosmosClient,
  makeCosmoshubPath,
  coins,
} from "@cosmjs/launchpad";

export default {
  components: { ...sp },
  computed: {},
  methods: {
    onClickCreate: async function() {
      let price = document.getElementById("price").value;
      let name = document.getElementById("name").value;
      let file = document.getElementById("file").files[0];
      let creator = document.getElementsByClassName("button__address")[0]
        .innerText;
      let base_req = { chain_id: "MusicChain", from: creator };
      let formData = new FormData();
      formData.append("price", price);
      formData.append("name", name);
      formData.append("file", file);
      formData.append("creator", creator);
      formData.append("base_req", JSON.stringify(base_req));
      const { data } = await axios.post(
        "http://localhost:1317/MusicChain/musics",
        formData
      );
      const { msg, fee, memo } = data.value;
      const mnemonic = localStorage.getItem("mnemonic");
      const wallet = await Secp256k1Wallet.fromMnemonic(
        mnemonic,
        makeCosmoshubPath(0),
        "cosmos"
      );
      const [{ address }] = await wallet.getAccounts();
      const client = new SigningCosmosClient(
        "http://localhost:1317",
        address,
        wallet
      );
      const res = await client.signAndPost(msg, fee, memo);
    },
  },
};
</script>
