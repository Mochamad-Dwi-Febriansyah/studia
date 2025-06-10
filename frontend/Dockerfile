FROM node:22-slim AS builder

WORKDIR /app

COPY package.json package-lock.json* ./

RUN npm install


COPY . ./

RUN npm run build

FROM node:22-slim

WORKDIR /app

ENV NODE_ENV=production

COPY --from=builder /app/.output  /app/.output

EXPOSE 3000

CMD ["node", ".output/server/index.mjs"]