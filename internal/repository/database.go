package repository

import (
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
)

type UserDatabaseRepo interface {
	SaveUser(user *entity.User) error
	CheckUserByEmail(email string) (*entity.User, error)
	StoreOTP(email, otp string, expiry time.Time) error
	VerifyOTP(email, otp string) (bool, error)
	GetUserByID(id string) (*entity.User, error)
	MyLearningCollection(myLearing *entity.MyLearning) error
	UpdatePassword(email, newHashedPassword string) error
}

type AdminDataBaseRepo interface {
	SaveAdmin(admin *entity.Admin) error
	CheckAdminByEmail(email string) *entity.Admin
}

type AboutSectionDatabaseRepo interface {
	CheckAboutSectionExists() (bool, error)
	CreateAboutSection(aboutSection *entity.AboutUsSection) error
	UpdateAboutSection(aboutSection *entity.AboutUsSection) error
	GetAboutSection() (*entity.AboutUsSection, error)
}

type PhilosopySectionDatabaseRepo interface {
	CheckPhilosopySectionExists() (bool, error)
	CreatePhilosopySection(philosopySection *entity.PhilosopySection) error
	UpdatePhilosopySection(philosopySection *entity.PhilosopySection) error
	GetPhilospySection() (*entity.PhilosopySection, error)
}

type StorySectionDatabaseRepo interface {
	CheckStorySectionExists() (bool, error)
	CreateStorySection(storysection *entity.StorySection) error
	UpdateStorySection(storySection *entity.StorySection) error
	GetstorySection() (*entity.StorySection, error)
}

type EventGallerySectionDatabaseRepo interface {
	AddEventGalleryImage(image *entity.EventGallary) error
	GetAllEventGalleryImages() ([]*entity.EventGallary, error)
	DeleteEventImagebyFileName(fileName string) error
}

type BlogsSectionDatabaseRepo interface {
	AddBlog(blog *entity.Blog) error
	GetAllBlogs() ([]*entity.Blog, error)
	DeleteBlogByFileName(fileName string) error
	GetBlogById(id string) (*entity.Blog, error)
}

type TestimonialDatabaseRepo interface {
	Addtestimonials(testimonials *entity.Testimonial) error
	GetVerifiedTestimonials() ([]*entity.Testimonial, error)
	GetUnverifiedTestimonials() ([]*entity.Testimonial, error)
	VerifyTestimonial(id string) error
	DeleteTestimonialByFileName(fileName string) error
}

type CourseDatabaseRepo interface {
	AddCourse(course *entity.Course) error
	GetAllTheCourses() ([]*entity.Course, error)
	DeleteCourseById(id string) error
	GetCourseByID(id string) (*entity.Course, error)
	GetPurchasedCoursesByUserId(userID string) ([]*entity.Course, error)
	GetCoursesExcludingUserPurchased(userID string) ([]*entity.Course, error)
}

type CartDatabaseRepo interface {
	AddToCart(item *entity.CartItem) error
	GetCartCourse(userID string) ([]*entity.Course, error)
	DeleteCartCourse(userID, courseID string) error
}

type PaymentDatabaseRepo interface {
	GetCourseAmountByIds(ids []string) (int32, error)
	AppendPurchasedCoursesIds(userID string, courseIDs []string) error
	SavePurchaseHistory(purchase *entity.PurchaseHistory) error
	GetSingleCoursePriceById(id string) (int32, error)
	DeleteCartCourseafterPayment(userID, courseID string) error
}

type PurchaseHistoryRepo interface {
	GetPurchaseHistoryByUserId(userID string) ([]*entity.PurchaseHistory, error)
}
