<template>
    <form
      class="message-input-form"
      @submit.prevent="sendMessage"
    >
      <input
        type="text"
        v-model="message"
        placeholder="메시지를 입력하세요"
      />
      <input
        type="submit"
        value="보내기"
        :disabled="!messageIsValid"
      />
    </form>
</template>

<script>
import { ref, computed } from "vue";

export default {
  name: "MessageInput",
  emits: ['send'],
  setup(props, { emit }) {
    const message = ref("");
    const sendMessage = () => {
        emit("send", message.value.trim())
        message.value = ''
    };
    const messageIsValid = computed(() => message.value.trim() !== "");
    return {
      message,
      sendMessage,
      messageIsValid
    };
  }
};
</script>

<style scoped>
.message-input-form {
  width: 500px;
  display: flex;
  bottom: 0;
  flex-grow: 0;
  height: 50px;
  background-color: white;
  border-radius: 10px;
  margin: 0 auto 20px auto;
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.15);
  overflow: hidden;
}

.message-input-form input[type="text"] {
  box-sizing: border-box;
  padding: 0px 20px;
  text-align: left;
  width: 80%;
  white-space: nowrap;
  font-size: 16px;
  line-height: 50px;
  height: 50px;
  margin: 0;
  float: left;
  border: none;
  width: 100%;
  color: #555;
  padding: 10px 20px;
}

.message-input-form input[type="text"]:focus {
  outline: none;
}

.message-input-form input[type="text"]::placeholder {
  color: #aaa;
}

.message-input-form input[type="submit"] {
  width: 100%;
  background-color: #ffed4e;
  color: #444;
  cursor: pointer;
  padding: 10px;
  font-size: 20px;
  box-sizing: border-box;
  border: 0;
  transition: background-color 0.2s ease-in-out;
  white-space: nowrap;
  font-size: 20px;
  width: 20%;
  padding: 0;
  line-height: 50px;
  height: 50px;
  float: right;
}

.message-input-form input[type="submit"]:focus {
  outline: none;
}
.message-input-form input[type="submit"]:hover {
  background-color: #e0d041;
}

.message-input-form input[type="submit"]:disabled {
  background-color: #ddd;
}
</style>