# 🌟 Moments - 极简朋友圈

[![release](https://img.shields.io/badge/release-更新记录-blue)](https://github.com/kingwrcy/moments/releases)
[![docker-release-status](https://img.shields.io/github/actions/workflow/status/kingwrcy/moments/docker-image-release.yml)](https://github.com/kingwrcy/moments/actions/workflows/docker-image-release.yml)
[![docker-pull](https://img.shields.io/docker/pulls/kingwrcy/moments)](https://hub.docker.com/repository/docker/kingwrcy/moments)
[![telegram-group](https://img.shields.io/badge/Telegram-group-blue)](https://t.me/simple_moments)
[![discussion](https://img.shields.io/badge/moments-论坛-blue)](https://discussion.mblog.club)

💡 **从 v0.2.1 开始，Moments 采用 Golang 重写服务端，包体积更小，功能更强！**  
🔍 **仍需 v0.2.0 版本？[点这里](https://github.com/kingwrcy/moments/blob/master/README.md)**

---

## 🎯 功能亮点

### 👤 用户系统

- 🏠 **默认账号**：`admin/a123456`，登录后可在后台修改
- 👥 **支持多用户模式**，可控制是否允许注册

### 📝 Memo 记录

- 🔖 **标签管理**，让内容更清晰
- 🖼️ **支持图片上传**，可存储至本地或 S3
- 🖼️ **自动生成缩略图**（目前仅支持本地上传）
- 📝 **支持 Markdown 语法**，后续增加更多标签支持
- ❤️ **点赞 & 评论**，可在后台控制评论功能
- 🎵 **支持网易云音乐、B 站视频、外部链接嵌入**
- 📖 **支持豆瓣读书 & 豆瓣电影引用**

### 🛠️ 其他功能

- 📱 **完美适配移动端**
- 🌙 **支持暗黑模式**
- ⏫ **回到顶部按钮**
- 🗄️ **数据库采用 SQLite**，随时可备份
- 🖼️ **支持自定义头图、头像、网站标题**

---

## 🚀 快速上手

### 🎯 一键启动 (推荐)

我们提供了便捷的一键启动脚本，支持开发模式和Docker模式：

```bash
# 开发模式 (前端 + 后端)
./start.sh dev

# Docker模式
./start.sh docker

# 查看所有选项
./start.sh help
```

📖 **详细使用指南**: [STARTUP_GUIDE.md](./STARTUP_GUIDE.md)

### 🛠️ 环境变量

Moments 支持以下 **环境变量** 进行配置：

| 变量名         | 说明                   | 默认值                                                                         |
| -------------- | ---------------------- | ------------------------------------------------------------------------------ |
| PORT           | 监听端口               | 37892                                                                          |
| CORS_ORIGIN    | 允许的跨域 Origin 列表 | 空，多个 Origin 可以使用英文逗号分隔，如 `http://127.0.0.1,http://10.10.10.10` |
| JWT_KEY        | JWT 密钥               | 空，不填写则随机生成，重启后需重新登录                                         |
| DB             | SQLite 数据库存放目录  | /app/data/db.sqlite                                                            |
| UPLOAD_DIR     | 上传文件本地目录       | /app/data/upload                                                               |
| LOG_LEVEL      | 日志级别               | info，可选 debug                                                               |
| ENABLE_SWAGGER | 启用 Swagger 文档      | false，可选 true，通过 `/swagger/index.html` 访问                              |

⚡ **支持 `.env` 文件加载环境变量**，示例：

```env
JWT_KEY=your_secret_key
LOG_LEVEL=info
```

---

## 🐳 使用 Docker

🔹 **启动容器**（需替换 `$JWT_KEY`）：

```bash
docker run -d \
  -e PORT=37892 \
  -e JWT_KEY=$JWT_KEY \
  -p 37892:37892 \
  -v ./moments:/app/data \
  --name moments \
  kingwrcy/moments:latest
```

📌 **持久化数据：** `/app/data` 挂载至 `./moments`  
📌 **可选：** `latest`（稳定版） 或 `dev`（开发版，功能前沿但相对不稳定）

### 📝 使用 Docker Compose

```yaml
services:
  moments:
    image: kingwrcy/moments:latest
    container_name: moments
    restart: always
    environment:
      PORT: 37892
      JWT_KEY: $JWT_KEY
    ports:
      - 37892:37892
    volumes:
      - ./moments:/app/data # 持久化数据到当前目录的 moments 文件夹
```

---

## 💻 可执行文件安装

🔽 **[下载最新版本](https://github.com/kingwrcy/moments/releases)**

> 示例（Windows 版）：

| 文件名                                         | 说明                         |
| ---------------------------------------------- | ---------------------------- |
| `moments-windows-amd64-x.x.x.exe.zip`          | **压缩包**，解压后可直接运行 |
| `moments-windows-amd64-x.x.x.exe-checksum.txt` | `MD5` 校验码，验证文件完整性 |

---

## 🔑 生成 JWT_KEY

📌 **方法 1：OpenSSL**

```bash
openssl rand -hex 32
```

📌 **方法 2：SHA256**

```bash
echo $RANDOM | sha256sum
```

📌 **方法 3：在线生成**（[点这里](https://tool.lu/uuid) 生成 UUID）

---

## 🛠️ 开发

### 🔧 依赖环境

📌 **后端：** `Go 1.23.3+`  
📌 **前端：** `NodeJS 18+`，推荐使用 `PNPM`  
📌 **VSCode 推荐插件：**

- `gitlens`（Git 扩展）
- `prettier`（代码格式化）
- `eslint`（代码规范检查）
- `golang`（Go 语言支持）

### 🏗️ 启动

#### 1️⃣ 使用 `make`（推荐）

后端：

```bash
cd moments
make backend-dev
```

前端（新终端）：

```bash
cd moments
make frontend-install
make frontend-dev
```

#### 2️⃣ 手动运行

后端：

```bash
cd moments/backend
go build -ldflags="-X main.version=local -X main.commitId=local" -o ./dist/moments
./dist/moments
```

前端：

```bash
cd moments/front
pnpm install
pnpm run dev
```

📍 启动后访问 `http://localhost:3000`

---

## 🌐 其他版本

| 项目                                                            | 演示地址                                                             |
| --------------------------------------------------------------- | -------------------------------------------------------------------- |
| [RandallAnjie/moments](https://github.com/RandallAnjie/moments) | [https://moments.randallanjie.com](https://moments.randallanjie.com) |

---

## ❤️ 致谢 Contributors

感谢所有贡献者！🎉

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/kingwrcy"><img src="https://avatars.githubusercontent.com/u/1247324?v=4?s=80" width="80px;" alt="kingwrcy"/><br /><sub><b>kingwrcy</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/RandallAnjie"><img src="https://avatars.githubusercontent.com/u/84122428?v=4?s=80" width="80px;" alt="Randall"/><br /><sub><b>Randall</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/Jonnyan404"><img src="https://avatars.githubusercontent.com/u/20352705?v=4?s=80" width="80px;" alt="jonny"/><br /><sub><b>jonny</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/akarikun"><img src="https://avatars.githubusercontent.com/u/11921182?v=4?s=80" width="80px;" alt="akari"/><br /><sub><b>akari</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/douseful"><img src="https://avatars.githubusercontent.com/u/52767905?v=4?s=80" width="80px;" alt="yee"/><br /><sub><b>yee</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://www.jschef.com"><img src="https://avatars.githubusercontent.com/u/38160059?v=4?s=80" width="80px;" alt="Chef"/><br /><sub><b>Chef</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://xwsir.cn"><img src="https://avatars.githubusercontent.com/u/17978673?v=4?s=80" width="80px;" alt="小王先森"/><br /><sub><b>小王先森</b></sub></a><br /></td>
    </tr>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://www.gooth.org"><img src="https://avatars.githubusercontent.com/u/126313?v=4?s=80" width="80px;" alt="Athurg Gooth"/><br /><sub><b>Athurg Gooth</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/xuewenG"><img src="https://avatars.githubusercontent.com/u/32838722?v=4?s=80" width="80px;" alt="xuewenG"/><br /><sub><b>xuewenG</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/Secretlovez"><img src="https://avatars.githubusercontent.com/u/40491055?v=4?s=80" width="80px;" alt="Secretlovez"/><br /><sub><b>Secretlovez</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/jkjoy"><img src="https://avatars.githubusercontent.com/u/23159890?v=4?s=80" width="80px;" alt="浪子"/><br /><sub><b>浪子</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/lateautumn2"><img src="https://avatars.githubusercontent.com/u/57248164?v=4?s=80" width="80px;" alt="lateautumn2"/><br /><sub><b>lateautumn2</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/Jinvic"><img src="https://avatars.githubusercontent.com/u/77521861?v=4?s=80" width="80px;" alt="Jinvic"/><br /><sub><b>Jinvic</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/dianso"><img src="https://avatars.githubusercontent.com/u/1454808?v=4?s=80" width="80px;" alt="DIANSO"/><br /><sub><b>DIANSO</b></sub></a><br /></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

📌 **欢迎贡献！** 详情见 [all-contributors](https://github.com/all-contributors/all-contributors) 规范。

---

## ⭐ Star History

[![Star History](https://api.star-history.com/svg?repos=kingwrcy/moments&type=Date)](https://star-history.com/#kingwrcy/moments&Date)

🔥 **如果你觉得 Moments 还不错，欢迎点个 Star！** 🚀
