FROM node:18-alpine AS build
WORKDIR /app
COPY frontend .
RUN npm install
RUN npm run build
RUN echo $(ls)

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
