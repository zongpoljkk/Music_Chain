<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form type="purchased" :fields="['MusicID', ]" module="MusicChain" />
		<sp-type-form type="musics" :fields="['mediaLink', 'price', 'name', ]" module="MusicChain" />
		<sp-type-form type="artist" :fields="['name', ]" module="MusicChain" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
