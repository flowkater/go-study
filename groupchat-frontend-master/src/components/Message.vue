<template>
  <div class="message" :class="{'my-message': mine}">
    <div class="sender" v-if="displaySender">
      {{message.sender}}
    </div>
    <div class="text">
      {{message.text}}
    </div>
  </div>
</template>

<script>
export default {
  name: "Message",
  props: {
    message: Object,
    mine: Boolean,
    displaySender: Boolean
  }
};
</script>

<style scoped>
.message {
  font-family: sans-serif;
  margin: 5px 0;
}

.sender {
  font-size: 12px;
  padding-left: 5px;
  padding-bottom: 5px;
}

.text {
  word-wrap: break-word;
  display: inline-block;
  margin: 0;
  flex-grow: 0;
  max-width: 300px;
  font-size: 14px;
  padding: 10px;
  border-radius: 10px;
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.15);
  background-color: #fff;
}

.my-message {
  margin-right: 0px;
  margin-left: auto;
}

.my-message .text {
  background-color: #ffed4e;
}
</style>