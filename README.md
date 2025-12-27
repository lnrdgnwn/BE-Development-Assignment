# BE : Soal 03

## Tech Stack
| Teknologi        | Kegunaan                                              |
| :--------------- | :---------------------------------------------------- |
| **Go (Golang)**  | Bahasa utama pengembangan backend.                    |
| **Fiber**        | Web framework cepat dan ringan.                       |
| **GORM**         | ORM untuk mempermudah interaksi dengan database.      |
| **MySQL**        | Database relasional untuk menyimpan data utama.       |
| **JWT (golang-jwt)** | Autentikasi berbasis token (stateless).          |
| **bcrypt**       | Hashing password dengan aman.                         |
| **dotenv**       | Manajemen environment variabel melalui file `.env`.   |
| **Swagger**      | Dokumentasi API otomatis dan interaktif.              |

## ⚙️ Instalasi & Konfigurasi

### 1️⃣ Clone Repository
```bash
git clone https://github.com/lnrdgnwn/BE-Development-Assignment.git
cd be-soal-03
```

2️⃣ Install Dependencies
```bash
go mod tidy
```

### 3️⃣ Ubah File .env.example ke .env dan lakukkan setupnya 
```bash
DB_HOST=your_db_host
DB_PORT=your_db_port
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
JWT_SECRET=your_jwt_secret
```

4️⃣ Jalankan Server
```bash
go run main.go
```
✔️Server akan berjalan di : 
```bash
👉 http://localhost:8080
```

## 🧩 Swagger API Documentation
Swagger digunakan untuk menampilkan dokumentasi API interaktif.
Jika sudah diaktifkan, kamu bisa membukanya melalui browser di alamat berikut:
```bash
👉 http://localhost:8080/swagger/index.html
```

## 🖥️ Cara Menjalankan Database (XAMPP)
1. Buka XAMPP Control Panel.
2. Jalankan Apache dan MySQL.
3. Buka phpMyAdmin melalui browser:
```bash
👉 http://localhost/phpmyadmin
```
4. Buat database baru dengan nama:
```bash
👉 your_db_name
```

## 📖 Dokumentasi API Endpoints 
### Authentication

#### 1. Register User

Mendaftarkan pengguna baru (CUSTOMER).

* **URL:** `/auth/register`
* **Method:** `POST`
* **Body (JSON):**
```json
{
    "name": "Leonardo Gunawan",
    "email": "leo@example.com",
    "password": "leo123456",
}

```

* **Response Success :**
```json
{
  "message": "User registered successfully"
}
```

Note : User otomatis memiliki role CUSTOMER ketika register, jika ingin menggunakan role ADMIN harus mengubah melalui DATABASE


#### 2. Login

Masuk untuk mendapatkan **JWT Token**. Token ini diperlukan untuk mengakses endpoint yang terkunci dan menandakan user telah terauentetikasi.

* **URL:** `/auth/login`
* **Method:** `POST`
* **Body (JSON):**
```json
{
    "email": "leo@gmail.com",
    "password": "leo12345"
}

```


* **Response Success:**
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9......."
}

```

### Users

#### 3. GetMyProfile

Mengambil data profile user yang sedang login.

* **URL:** `/users/me`
* **Method:** `GET`
* **Auth:** *Required*
* * **Response Example:**
```json
{
    "id": 1,
    "name": "Admin",
    "email": "admin@gmail.com",
    "role": "ADMIN"
    "created_at": "2025-12-27T12:37:20.663+07:00",
    "updated_at": "2025-12-27T12:37:20.663+07:00"
}

```


#### 4. Get All Users

Mengambil semua data user (Admin only).

* **URL:** `/users/me`
* **Method:** `GET`
* **Auth:** *Required(ADMIN)*
* * **Response Example:**
```json
[
  {
    "id": 1,
    "name": "Admin",
    "email": "admin@gmail.com",
    "role": "ADMIN"
    "created_at": "2025-12-27T12:37:20.663+07:00",
    "updated_at": "2025-12-27T12:37:20.663+07:00"
  },
  {
    "id": 2,
    "name": "User",
    "email": "user@mgail.com",
    "role": "CUSTOMER"
    "created_at": "2025-12-27T10:31:39.048+07:00",
    "updated_at": "2025-12-27T10:31:39.048+07:00"
  }
]

```



### Events

#### 5. Get All Events

Mengambil daftar semua event yang tersedia.

* **URL:** `/events`
* **Method:** `GET`
* **Auth:** *Public*
* **Response Example:**
```json
[
    {
        "id": 1,
        "title": "Seminar A",
        "location": "JL Mawar 15",
        "total_ticket": 10
        ...
    } ,
    {
        "id": 2,
        "title": "Konser Musik b",
        "location": "Jl Garuda 1",
        "total_ticket": 10
        ...
    } ,
]

```

#### 6. Get Events By Id

Mengambil data satu event yang tersedia.

* **URL:** `/events/:id`
* **Method:** `GET`
* **Auth:** *Public*
* **Response Example:**
```json
[
    {
        "id": 1,
        "title": "Seminar A",
        "location": "JL Mawar 15",
        "total_ticket": 10
        ...
    }
]

```

#### 7. Create Event (Admin Only)

Membuat event baru yang hanya dapat diakses oleh user dengan role `ADMIN`.

* **URL:** `/events`
* **Method:** `POST`
* **Headers:**
* **Auth:** *Required(ADMIN)*
* **Body (JSON):**
```json

{
  "id": 1,
  "title": "Seminar A",
  "description": "Pengembangan diri",
  "event_date": "2025-12-27T10:31:39.048+07:00",
  "location": "Jl Garuda 1",
  "total_ticket": 30,
}

```
* **Response Example:**
```json
{
  "id": 1,
  "title": "Seminar A",
  "description": "Pengembangan diri",
  "event_date": "2025-12-27T10:31:39.048+07:00",
  "location": "Jl Garuda 1",
  "total_ticket": 30,
  "available_ticket": 30,
  "organizer_id": 4,
  "status": "PUBLISHED",
  "created_at": "2025-12-27T18:27:19.5620069+07:00",
  "updated_at": "2025-12-27T18:27:19.5620069+07:00",
  "organizer": {
      "id": 4,
      "name": "Leonardo Gunawan",
      "email": "leo@gmail.com",
      "role": "ADMIN",
      "created_at": "2025-12-27T12:37:20.663+07:00",
      "updated_at": "2025-12-27T12:37:20.663+07:00"
  }
}

```

#### 8. Update Event (Admin Only)

Mengupdate event yang hanya dapat dilakukan oleh user dengan role `ADMIN`.

* **URL:** `/events/:id`
* **Method:** `PUT`
* **Headers:**
* **Auth:** *Required(ADMIN)*
* **Body (JSON):**
```json

{
  "title": "Seminar A",
  "description" :"Pengembangan diri Gen Z"
  "status": "FINISHED",
  "event_date": "2025-14-27T10:31:39.048+07:00",
  "location": "Jl Bangau 1",
  "total_ticket": 40,
}

```
* **Response Example:**
```json
{
  "id": 1,
  "title": "Seminar A",
  "description" :"Pengembangan diri Gen Z"
  "status": "FINISHED",
  "event_date": "2025-14-27T10:31:39.048+07:00",
  "location": "Jl Bangau 1",
  "total_ticket": 40,
  "available_ticket": 30,
  "status": "PUBLISHED",
  "created_at": "2025-12-27T18:27:19.5620069+07:00",
  "updated_at": "2025-12-27T18:27:19.5620069+07:00",
  "organizer": {
      "id": 4,
      "name": "Leonardo Gunawan",
      "email": "leo@gmail.com",
      "role": "ADMIN",
      "created_at": "2025-12-27T12:37:20.663+07:00",
      "updated_at": "2025-12-27T12:37:20.663+07:00"
  }
}

```


### Transactions

####  9. Get My Transactions
Mengambil semua transaksi user yang sedang login.
* **URL:** `/transactions`
* **Method:** `GET`
* **Auth:** *Required*
* **Response Example:**
```json
[
  {
    "id": 1,
    "user_id": 4,
    "event_id": 1,
    "quantity": 5,
    "status": "SUCCESS",
    "created_at": "2025-12-27T19:47:45.043+07:00",
    "updated_at": "2025-12-27T19:47:45.043+07:00",
    "user": {
      "id": 0,
      "name": "",
      "email": "",
      "role": "",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z"
    },
    "event": {
      "id": 1,
      "title": "Seminar A",
      "description": "Pengembangan diri",
      "event_date": "2025-12-27T10:31:39.048+07:00",
      "location": "Jl Garuda 1",
      "total_ticket": 30,
      "available_ticket": 18,
      "status": "PUBLISHED",
      "created_at": "2025-12-27T18:27:19.562+07:00",
      "updated_at": "2025-12-27T20:29:22.467+07:00",
      "organizer": {
        "id": 0,
        "name": "",
        "email": "",
        "role": "",
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z"
      }
    }
  },
  {
    "id": 2,
    "user_id": 4,
    "event_id": 1,
    "quantity": 3,
    "status": "SUCCESS",
    "created_at": "2025-12-27T20:16:46.87+07:00",
    "updated_at": "2025-12-27T20:16:46.87+07:00",
    "user": {
      "id": 0,
      "name": "",
      "email": "",
      "role": "",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z"
    },
    "event": {
      "id": 1,
      "title": "Seminar A",
      "description": "Pengembangan diri",
      "event_date": "2025-12-27T10:31:39.048+07:00",
      "location": "Jl Garuda 1",
      "total_ticket": 30,
      "available_ticket": 18,
      "status": "PUBLISHED",
      "created_at": "2025-12-27T18:27:19.562+07:00",
      "updated_at": "2025-12-27T20:29:22.467+07:00",
      "organizer": {
        "id": 0,
        "name": "",
        "email": "",
        "role": "",
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z"
      }
    }
  }
]
``` 

####  10. Book Ticket 

Membeli tiket untuk event tertentu dengan membatasi jumlah beli harus kurang dari sama dengan available ticket

* **URL:** `/transactions`
* **Method:** `POST`
* **Auth:** *Required*

* **Body (JSON):**
```json
{
    "event_id": 1,
    "quantity": 2
}

```


* **Response Success:**
```json
{
    {
  "id": 4,
  "user_id": 4,
  "event_id": 1,
  "quantity": 1,
  "status": "SUCCESS",
  "created_at": "2025-12-27T20:29:22.468+07:00",
  "updated_at": "2025-12-27T20:29:22.468+07:00",
  "user": {
    "id": 0,
    "name": "",
    "email": "",
    "role": "",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
  "event": {
    "id": 1,
    "title": "Seminar A",
    "description": "Pengembangan diri",
    "event_date": "2025-12-27T10:31:39.048+07:00",
    "location": "Jl Garuda 1",
    "total_ticket": 30,
    "available_ticket": 18,
    "status": "PUBLISHED",
    "created_at": "2025-12-27T18:27:19.562+07:00",
    "updated_at": "2025-12-27T20:29:22.467+07:00",
    "organizer": {
      "id": 0,
      "name": "",
      "email": "",
      "role": "",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z"
    }
  }
}
}
```

* **Response Error (Jika Memesan tiket Lebih dari available_ticket):**
```json
{
  "error": "Not enough tickets available"
}

```


## Struktur Project

```
be-soal-03/
├── config/
│   ├── env.go   # Koneksi Database MySQL 
├── controllers/
│   ├── auth_controller.go        # Login & Register Logic
│   ├── event_controller.go       # CRUD Event
│   └── transaction_controller.go # Pembelian Tiket
├── database/
│   ├── database.go 
├── docs/
│   ├── docs.go  
│   ├── swagger.json
│   └── swagger.yaml
├── middlewares/
│   ├── auth_middlewares.go    
│   └── admin_middlewares.go       
├── models/
│   ├── event.go
│   ├── transaction.go
│   └── user.go
├── routes/
│   ├── auth_routes.go
│   ├── transaction_routes.go
│   ├── event_routes.go
│   ├── transaction_routes.go
│   └── routes.go
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── link.txt
└── main.go

```
