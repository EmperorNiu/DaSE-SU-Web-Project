// module.exports = {
//   presets: [
//     '@vue/cli-plugin-babel/preset'
//   ],
//   plugins: [
//     [
//       'component',
//       {
//         libraryName: 'element-ui',
//         styleLibraryName: 'theme-chalk'
//       }
//     ]
//   ]
// }
module.exports = {
  presets: [
    '@vue/cli-plugin-babel/preset'
  ],
  plugins: [
    [
      'component',
      {
        libraryName: 'mint-ui',
        style: true
      },
      'mint-ui'
    ],
    [
      'component',
      {
        libraryName: 'element-ui',
        styleLibraryName: 'theme-chalk',
        style: true
      },
      'element-ui'
    ]
  ]
}
