<template>
 <div class="menubar">
      <div class="left">
        <div
          class="back-button"
          @click="$emit('disconnect')"
        ><svg
            width="1em"
            height="1em"
            viewBox="0 0 16 16"
            class="bi bi-arrow-left"
            fill="currentColor"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fill-rule="evenodd"
              d="M15 8a.5.5 0 0 0-.5-.5H2.707l3.147-3.146a.5.5 0 1 0-.708-.708l-4 4a.5.5 0 0 0 0 .708l4 4a.5.5 0 0 0 .708-.708L2.707 8.5H14.5A.5.5 0 0 0 15 8z"
            />
          </svg>뒤로</div>
      </div>
      <div class="room-name">{{roomName}}</div>
      <div class="right"></div>
    </div>
</template>

<script>
export default {
    name: 'MenuBar',
    emits: ['disconnect'],
    props: {
        roomName: String
    }
}
</script>

<style scoped>
.menubar {
  font-family: sans-serif;
  height: 60px;
  width: 100%;
  background-color: rgba(255, 255, 255, 0.6);
  position: fixed;
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.15);
  backdrop-filter: saturate(180%) blur(20px);
  display: flex;
  justify-content: center;
}
.menubar .room-name {
  line-height: 60px;
  text-align: center;
  color: #444;
  font-size: 26px;
}

.menubar .left {
  flex: 1;
}

.menubar .back-button {
  line-height: 60px;
  text-align: center;
  width: 100px;
  color: #444;
  font-size: 16px;
  transition: background-color 0.2s ease-in-out;
}

.menubar .back-button:hover {
  background-color: rgba(0, 0, 0, 0.15);
  cursor: pointer;
}
.menubar .back-button svg {
  padding: 0 8px;
}

.menubar .right {
  flex: 1;
}
</style>