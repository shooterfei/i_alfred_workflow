import typescript from "@rollup/plugin-typescript"
import { nodeResolve } from "@rollup/plugin-node-resolve"
import commonjs from "@rollup/plugin-commonjs"
import nodeResolveDefaultImport from "@rollup/plugin-node-resolve"
import json from "@rollup/plugin-json"

export default {
  input: './main.ts',//入口文件
  output: {
    // dir: 'dist',
    file: 'dist/bundle.mjs',//打包后的存放文件, 后缀名可根据需要进行修改, js cjs mjs
    format: 'es',//输出格式 amd es6 iife umd cjs
    inlineDynamicImports: true,
    name: 'bundleName'//如果iife,umd需要指定一个全局变量
  },
  external: [ 'canvas' ],
  plugins: [
    typescript({
      tsconfig: "./tsconfig.json"
    }),
    commonjs({ extensions: [".js", ".ts", ".d.ts", ".node"] }),
    nodeResolve(),
    json(),
    nodeResolveDefaultImport()
  ]
}
