package routes

import (
	"github.com/bhushan-aruto/aspiration-matters-backend/config"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/messaging"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/server/handlers"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/storage"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

func InitRoutes(
	cfg *config.Config,
	e *echo.Echo,
	db *database.MongoDBdatabase,
	producer *messaging.RabbitMQProducer,
	s3Conn *storage.S3Connection,
	paymentGatewayRepo repository.PaymentGateWayRepo,
) {

	userHandler := handlers.NewUserHandler(db, producer)
	adminHandler := handlers.NewAdminHandler(db)

	aboutSectionHandler := handlers.NewAboutSectionHandler(
		cfg.S3BucketName,
		cfg.AboutSectionDirName,
		cfg.CloudFrontUrl,
		db,
		s3Conn,
	)

	philosopySectionHandler := handlers.NewPhilosopySectionHandler(
		cfg.S3BucketName,
		cfg.PhilosopysectionDirName,
		cfg.CloudFrontUrl,
		db,
		s3Conn,
	)
	storySectionHandler := handlers.NewStorySectionHandler(
		cfg.S3BucketName,
		cfg.StorysectionDirName,
		cfg.CloudFrontUrl,
		db,
		s3Conn,
	)
	eventGalleryHandler := handlers.NewEventGalleryHandler(
		cfg.S3BucketName,
		cfg.EventGallerysectionDirName,
		cfg.CloudFrontUrl,
		db,
		s3Conn,
	)
	blogsHandler := handlers.NewBlogsHandler(
		cfg.S3BucketName,
		cfg.BlogsSectionDirName,
		cfg.CloudFrontUrl,
		db,
		s3Conn,
	)
	testimonialHandler := handlers.NewTestimonialsHandler(
		cfg.S3BucketName,
		cfg.TestimonialSectionDirName,
		cfg.CloudFrontUrl,
		db,
		s3Conn,
	)
	courseHandler := handlers.NewCourseHandler(
		cfg.CourseS3BucketName,
		cfg.CourseSectionDirName,
		cfg.CourseCloufontUrl,
		db,
		s3Conn,
	)
	cartHandler := handlers.NewCartHandler(db)

	paymentHandler := handlers.NewPaymentHandler(
		paymentGatewayRepo,
		db,
	)

	contactUsecase := usecase.NewContactUs()
	contactHandler := handlers.NewContactHandler(contactUsecase, cfg)

	purchaseHandler := handlers.NewPurchaseHistoryHandler(db)

	//user routes
	userGroup := e.Group("/user")
	userGroup.POST("/signup", userHandler.UserSignupHandler)
	userGroup.POST("/send-otp", userHandler.SentOTPHandler)
	userGroup.POST("/verify-otp", userHandler.VerifyAndLoginHandler)
	userGroup.GET("/:id", userHandler.GetUserByIdHandler)

	userGroup.POST("/forgot-password", userHandler.ForgotPasswordSendOTPHandler)
	userGroup.POST("/reset-password", userHandler.ResetPasswordHandler)

	//adimn routes
	adminGroup := e.Group("/admin")
	adminGroup.POST("/signup", adminHandler.AdminSignupHandler)
	adminGroup.POST("/login", adminHandler.AdminLoginHandler)

	//about scetion routes
	aboutGroup := e.Group("/about")
	aboutGroup.GET("", aboutSectionHandler.GetAboutSectionHandler)
	aboutGroup.POST("/create", aboutSectionHandler.CreateAboutSectionHandler)
	aboutGroup.PUT("/update/image1", aboutSectionHandler.UpdateAbouSectionImage1Handler)
	aboutGroup.PUT("/update/image2", aboutSectionHandler.UpdateAboutSectionImage2Handler)
	aboutGroup.DELETE("/delete/image1/:fileName", aboutSectionHandler.DeleteAboutSectionImage1Handler)
	aboutGroup.DELETE("/delete/image2/:fileName", aboutSectionHandler.DeleteAboutSectionImage2Handler)

	//philosopy scetion routes
	philosopyGroup := e.Group("/philosopy")
	philosopyGroup.GET("", philosopySectionHandler.GetPhilosopysectionHandler)
	philosopyGroup.POST("/create", philosopySectionHandler.CreatePhilosopySectionHandler)
	philosopyGroup.PUT("/update/image1", philosopySectionHandler.UpdatePhilosopySectionImage1Handler)
	philosopyGroup.PUT("/update/image2", philosopySectionHandler.UpdatePhilosopySectionImage2Handler)
	philosopyGroup.DELETE("/delete/image1/:fileName", philosopySectionHandler.DeletePhilosopySectionImage1Handler)
	philosopyGroup.DELETE("/delete/image2/:fileName", philosopySectionHandler.DeletePhilosopySectionImage2Handler)

	//story scetion routes
	storyGroup := e.Group("/story")
	storyGroup.GET("", storySectionHandler.GetStorySectionHandler)
	storyGroup.POST("/create", storySectionHandler.CreateStorySectionHandler)
	storyGroup.PUT("/update/image1", storySectionHandler.UpdateStorySectionImage1Handler)
	storyGroup.PUT("/update/image2", storySectionHandler.UpdateStorySectionImage2Handler)
	storyGroup.PUT("/update/image3", storySectionHandler.UpdateStorySectionImage3Handler)
	storyGroup.PUT("/update/image4", storySectionHandler.UpdateStorySectionImage4Handler)
	storyGroup.DELETE("/delete/image1/:fileName", storySectionHandler.DeleteStorySectionImage1Handler)
	storyGroup.DELETE("/delete/image2/:fileName", storySectionHandler.DeleteStorySectionImage2Handler)
	storyGroup.DELETE("/delete/image3/:fileName", storySectionHandler.DeleteStorySectionImage3Handler)
	storyGroup.DELETE("/delete/image4/:fileName", storySectionHandler.DeleteStorySectionImage4Handler)

	//event-gallery scetion routes
	eventGalleryGroup := e.Group("/eventgallery")
	eventGalleryGroup.POST("/", eventGalleryHandler.EventGalleryUploadImageHandler)
	eventGalleryGroup.GET("/", eventGalleryHandler.GetAllEventGalleryImagesHandler)
	eventGalleryGroup.DELETE("/:fileName", eventGalleryHandler.DeleteImageHandler)

	//blogs section routes
	blogGroup := e.Group("/blog")
	blogGroup.POST("/", blogsHandler.UploadBloghandler)
	blogGroup.GET("/", blogsHandler.GetAllBlogsHandler)
	blogGroup.DELETE("/:fileName", blogsHandler.DeletBlogsHandler)
	blogGroup.GET("/:id", blogsHandler.GetBlogHandler)

	//blogs section routes
	testimonialGroup := e.Group("/testimonial")
	testimonialGroup.POST("/", testimonialHandler.AddTestimonialsHandler)
	testimonialGroup.GET("/verified", testimonialHandler.GetVerifiedtestimonials)
	testimonialGroup.GET("/unverified", testimonialHandler.GetUnVerifiedtestimonials)
	testimonialGroup.PUT("/approve/:id", testimonialHandler.ApproveTestimonials)
	testimonialGroup.DELETE("/:fileName", testimonialHandler.DeleteTestimonial)

	//course admin section
	course := e.Group("/course")
	course.POST("/upload", courseHandler.UploadCourseHandler)
	course.GET("/get", courseHandler.GetAllCourseUsecase)
	course.DELETE("/delete/:id", courseHandler.DeleteCourseByIdHandler)
	course.GET("/:id", courseHandler.GetCourseByIdHandler)
	course.GET("/purchased/:user_id", courseHandler.GetPurchasedCoursesByUserIdHandler)
	course.GET("/user/:user_id", courseHandler.GetAllCourseUsecaseExpectPurchasedHandler)

	//cart
	cartGroup := e.Group("/cart")
	cartGroup.POST("/add", cartHandler.AddToCartHandler)
	cartGroup.GET("/course", cartHandler.GetAllCartCourseHandler)
	cartGroup.DELETE("/:user_id/:course_id", cartHandler.DeleteCartCourseHandler)

	//contact
	contactGroup := e.Group("/contact")
	contactGroup.POST("/email", contactHandler.HandleEmailContact)
	contactGroup.POST("/whatsapp", contactHandler.HandleWhatsAppContact)
	contactGroup.POST("/course-enquiry", contactHandler.CourseEnqiryHandler)

	//payment gateway
	paymentGroup := e.Group("/payment")
	paymentGroup.POST("/order", paymentHandler.CreateOrderHandler)
	paymentGroup.POST("/verify", paymentHandler.VerifyPaymentHandler)

	//purchase HIstory
	purchaseGroup := e.Group("/purchase")
	purchaseGroup.GET("/:user_id", purchaseHandler.GetPurchaseHistoryByUserIdHandler)

}
