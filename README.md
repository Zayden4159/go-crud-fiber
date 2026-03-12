# CRUD WITH FIBER พื้นฐาน

ตัวอย่าง REST API พื้นฐานด้วย Go + Fiber ครอบคลุม CRUD และการอัพโหลดไฟล์

## Features
- ดึงข้อมูล User ทั้งหมด
- ดึงข้อมูล User รายคน
- สร้าง User
- แก้ไขข้อมูล User
- ลบ User
- อัพโหลดไฟล์

## Setup
```bash
git clone https://github.com/yourusername/crud-fiber.git
cd crud-fiber
go mod tidy
go run .
```

## API  Endpoints
| Method | Endpoint | Description |
|---|---|---|
| GET | /users | ดึง User ทั้งหมด |
| GET | /user/:id | ดึง User รายคน |
| POST | /user | สร้าง User |
| PUT | /user/:id | แก้ไขข้อมูล User |
| DELETE | /user/:id | ลบ User |
| POST | /upload | อัพโหลดไฟล์ |

## Tech Stack
- Go
- gofiber

## License
MIT