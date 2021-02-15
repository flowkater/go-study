<template>
  <div class="auth">
    <form @submit.prevent="submit">
      <input
        type="text"
        placeholder="닉네임"
        v-model="username"
      />
      <input
        type="password"
        placeholder="비밀번호"
        v-model="password"
      />
      <button
        :disabled="!valid">{{mode === 'signin' ? '로그인하기' : '가입하기'}}</button>
      <a
        href="#"
        @click.prevent="toggleMode"
      >
        {{mode == 'signin' ? '가입하러 가기' : '로그인하러 가기'}}
      </a>
    </form>
  </div>
</template>

<script>
import { computed, ref } from 'vue'
export default {
  name: 'Auth',
  emits: ['token'],
  setup(props, { emit }) {
    const username = ref('')
    const password = ref('')
    const mode = ref('signin')

    return {
      username,
      password,
      mode,
      toggleMode: () => {
        if (mode.value === 'signin') mode.value = 'signup'
        else mode.value = 'signin'
      },
      valid: computed(() => username.value.trim() !== '' && password.value !== ''),
      submit: async () => {
        try {
          const resp = await fetch(`http://localhost:8080/${mode.value}`,
            {
              method: 'POST',
              body: JSON.stringify({
                username: username.value.trim(),
                password: password.value
              })
            })
          if (resp.status !== 200) {
            throw 'expected 200'
          }
          const parsed = await resp.json()
          emit('token', parsed.token)
        } catch {
          alert(mode.value + ' error')
        }
      }
    }
  }
}
</script>

<style scoped>
.auth {
  box-sizing: border-box;
  width: 100%;
  font-family: sans-serif;
}

a {
    font-size: 12px;
}

input {
  box-sizing: border-box;
  margin: 10px 0;
  width: 100%;
  padding: 2px 5px;
  display: block;
}
button {
  box-sizing: border-box;
  width: 100%;
  display: block;
  margin: auto;
}
</style>