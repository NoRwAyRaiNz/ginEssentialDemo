package controller

import (
	"errors"
	"ginEssential/Model"
	"ginEssential/common"
	"ginEssential/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ICategoryController interface {
	RestController //封装增删改查代码 以便复用
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(Model.Category{})

	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory Model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
	} else {
		c.DB.Create(&requestCategory)

		response.Success(ctx, gin.H{"category": requestCategory}, "")
	}

}

func (c CategoryController) Update(ctx *gin.Context) {

	// 绑定body中的参数
	var requestCategory Model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
	}

	// 获取path中饿的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory Model.Category
	err := c.DB.First(&updateCategory, categoryId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Fail(ctx, "分类不存在", nil)

	} else {
		//更新分类
		// map
		// struct
		// name value
		c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

		response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
	}
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中饿的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var category Model.Category
	err := c.DB.First(&category, categoryId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Fail(ctx, "分类不存在", nil)

	} else {
		response.Success(ctx, gin.H{"category": category}, "")
	}

}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中饿的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.DB.Delete(Model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, "删除失败，请重试", nil)
		return

	} else {
		response.Success(ctx, nil, "")
	}

}
