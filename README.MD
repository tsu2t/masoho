## masoho

シンプルな **Go 言語製 REST API サーバ** のサンプルプロジェクトです。  
MASOHO（真朱）は万葉集に登場する、日本の伝統色です。まそほの他、「しんしゅ」「まそお」とも読みます。カラーコードは#D9341Dです。
勉強用に極力Cursorへの指示のみで実装していきます。

---

## 動作環境

- **Go**: 1.23 以降（`go.mod` に準拠）
- OS: macOS / Linux / Windows など、Go が動作する環境

---

## セットアップ

```bash
git clone <このリポジトリのURL>
cd masoho/backend
# 依存関係取得（標準ライブラリのみだが念のため）
go mod tidy
cd ../frontend
npm install   # or pnpm install / yarn install
```

---

## 起動方法

```bash
cd masoho
go run ./backend
# → http://localhost:8080 で REST API が待ち受け

cd masoho/frontend
npm run dev
# → 通常 http://localhost:5173 でフロントが起動
```

---

## ライセンス
MIT

