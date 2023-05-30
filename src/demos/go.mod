// 模块名
module demos

// go sdk 版本
go 1.20

// 当前module依赖的包
require (
	//dependency latest
)

// 排除第三方包
exclude (
	//dependency latest
)

// 修改依赖包的路径或版本(依赖包迁移，原始包无法访问，本地包替换原始包等场景)
replace (
	//source latest => target latest
)

// 撤回 (当前module作为其他module的依赖，某个版本出现问题需要撤回的场景)
retract {}
