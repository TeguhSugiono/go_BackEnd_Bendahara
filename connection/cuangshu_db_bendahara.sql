{
	"info": {
		"_postman_id": "1058ea6e-621f-48d3-940d-2d196bb1d8d0",
		"name": "Rest API Bendahara",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "2494232"
	},
	"item": [
		{
			"name": "1.Master User",
			"item": [
				{
					"name": "Daftar User",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"username\" : \"teguh\",\r\n    \"password\" : \"12345\",\r\n    \"full_name\" : \"Teguh Sugiono\"\r\n}\r\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/users/signup",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"users",
								"signup"
							]
						}
					},
					"response": []
				},
				{
					"name": "Login User",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"username\" : \"teguh\" ,\r\n    \"password\" : \"12345\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://10.5.57.55:2022/api/v1/users/login",
							"protocol": "http",
							"host": [
								"10",
								"5",
								"57",
								"55"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"users",
								"login"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "2.Master Jenis Transaksi",
			"item": [
				{
					"name": "Tampil Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterjenistrans/showjenistrans?page=1&perpage=10&search=na&sort=asc",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterjenistrans",
								"showjenistrans"
							],
							"query": [
								{
									"key": "page",
									"value": "1"
								},
								{
									"key": "perpage",
									"value": "10"
								},
								{
									"key": "search",
									"value": "na"
								},
								{
									"key": "sort",
									"value": "asc"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "List Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterjenistrans/listjenistrans",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterjenistrans",
								"listjenistrans"
							]
						}
					},
					"response": []
				},
				{
					"name": "Create Data",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"proses_uang\" : \"Uang Keluar\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterjenistrans/insertjenistrans",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterjenistrans",
								"insertjenistrans"
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Proses_uang\" : \"Uang Panas\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterjenistrans/updatejenistrans/3",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterjenistrans",
								"updatejenistrans",
								"3"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterjenistrans/deletejenistrans/2",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterjenistrans",
								"deletejenistrans",
								"2"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "3.Master Group Kategori",
			"item": [
				{
					"name": "Tampil Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/mastergroupkategori/showgroupkategori?page=1&perpage=10",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"mastergroupkategori",
								"showgroupkategori"
							],
							"query": [
								{
									"key": "page",
									"value": "1"
								},
								{
									"key": "perpage",
									"value": "10"
								},
								{
									"key": "search",
									"value": "gaji",
									"disabled": true
								},
								{
									"key": "sort",
									"value": "asc",
									"disabled": true
								},
								{
									"key": "filter",
									"value": "Uang Masuk",
									"disabled": true
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "List Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/mastergroupkategori/listgroupkategori",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"mastergroupkategori",
								"listgroupkategori"
							]
						}
					},
					"response": []
				},
				{
					"name": "Create Data",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_jenis\" : 1,\r\n    \"nm_group\" : \"Uang Masuk Eksternal\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/mastergroupkategori/insertgroupkategori",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"mastergroupkategori",
								"insertgroupkategori"
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"nm_group\" : \"Biaya Makan Anak Buaya\" ,\r\n    \"Kd_jenis\" : 1 \r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/mastergroupkategori/updategroupkategori/10",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"mastergroupkategori",
								"updategroupkategori",
								"10"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/mastergroupkategori/deletegroupkategori/10",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"mastergroupkategori",
								"deletegroupkategori",
								"10"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "4.Master Kategori Uang",
			"item": [
				{
					"name": "Tampil Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterkategoriuang/showkategoriuang?page=1&perpage=100",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterkategoriuang",
								"showkategoriuang"
							],
							"query": [
								{
									"key": "page",
									"value": "1"
								},
								{
									"key": "perpage",
									"value": "100"
								},
								{
									"key": "search",
									"value": "Terang",
									"disabled": true
								},
								{
									"key": "sort",
									"value": "asc",
									"disabled": true
								},
								{
									"key": "filter",
									"value": "Biaya Pembayaran Daftar Ulang",
									"disabled": true
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "List Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterkategoriuang/listkategoriuang",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterkategoriuang",
								"listkategoriuang"
							]
						}
					},
					"response": []
				},
				{
					"name": "Create Data",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 10,\r\n    \"nm_kategori\" : \"Dana Masuk dari Bos (Pemilik Yayasan)\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterkategoriuang/insertkategoriuang",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterkategoriuang",
								"insertkategoriuang"
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 1 ,\r\n    \"nm_kategori\" :  \"TEST INPUT DATA\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterkategoriuang/updatekategoriuang/15",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterkategoriuang",
								"updatekategoriuang",
								"15"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterkategoriuang/deletekategoriuang/15",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterkategoriuang",
								"deletekategoriuang",
								"15"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "5.Master Kategori Sub Uang",
			"item": [
				{
					"name": "Tampil Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterkategoriuang/showsubkategoriuang?page=1&perpage=10",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterkategoriuang",
								"showsubkategoriuang"
							],
							"query": [
								{
									"key": "page",
									"value": "1"
								},
								{
									"key": "perpage",
									"value": "10"
								},
								{
									"key": "search",
									"value": "Terang",
									"disabled": true
								},
								{
									"key": "sort",
									"value": "asc",
									"disabled": true
								},
								{
									"key": "filter",
									"value": "PPDB",
									"disabled": true
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "List Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterkategoriuang/listsubkategoriuang",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterkategoriuang",
								"listsubkategoriuang"
							]
						}
					},
					"response": []
				},
				{
					"name": "Create Data",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Kd_kategori\" : 1,\r\n    \"Nm_sub_kategori\" : \"TEST Inputc\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterkategoriuang/insertsubkategoriuang",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterkategoriuang",
								"insertsubkategoriuang"
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_kategori\" : 1 ,\r\n    \"nm_sub_kategori\" :  \"TEST INPUT\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterkategoriuang/updatesubkategoriuang/4",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterkategoriuang",
								"updatesubkategoriuang",
								"4"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterkategoriuang/deletesubkategoriuang/4",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterkategoriuang",
								"deletesubkategoriuang",
								"4"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "6.Master Configurasi Periode",
			"item": [
				{
					"name": "Tampil Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/settingperiode/showconfperiode?page=1&perpage=10",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"settingperiode",
								"showconfperiode"
							],
							"query": [
								{
									"key": "page",
									"value": "1"
								},
								{
									"key": "perpage",
									"value": "10"
								},
								{
									"key": "search",
									"value": "2023",
									"disabled": true
								},
								{
									"key": "sort",
									"value": "asc",
									"disabled": true
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "List Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNDg5NjY4fQ.rze0BLyiqiT5OLVIWnE_VVBHSHB8vkE34jsd84XRvBY",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/settingperiode/listconfperiode",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"settingperiode",
								"listconfperiode"
							]
						}
					},
					"response": []
				},
				{
					"name": "Create Data",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tahun_akademik\" : \"2022/2023\",\r\n    \"nm_kelas\" : \"XII\",\r\n    \"biaya_spp\" : 750000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/settingperiode/insertconfperiode",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"settingperiode",
								"insertconfperiode"
							]
						}
					},
					"response": []
				},
				{
					"name": "XEdit Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNDg5NjY4fQ.rze0BLyiqiT5OLVIWnE_VVBHSHB8vkE34jsd84XRvBY",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"biaya_spp\" : 500000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/settingperiode/updateconfperiode/4",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"settingperiode",
								"updateconfperiode",
								"4"
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit Data All",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNDg5NjY4fQ.rze0BLyiqiT5OLVIWnE_VVBHSHB8vkE34jsd84XRvBY",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tahun_akademik\" : \"2022-2023\",\r\n    \"nm_kelas\" : \"X\",\r\n    \"biaya_spp\" : 125000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/settingperiode/updateconfperiodeall",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"settingperiode",
								"updateconfperiodeall"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Data",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n     \"tahun_akademik\" : \"2022-2023\",\r\n     \"nm_kelas\" : \"X\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/settingperiode/deleteconfperiode",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"settingperiode",
								"deleteconfperiode"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "7.Master Tahun Akademik",
			"item": [
				{
					"name": "List Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyMzY4Mjk4fQ.ru0YaAgu76dGwSK0adwrcDwuIcBOgX6TJUbNFIhCz4A",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/akademik/listtahunakademik",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"akademik",
								"listtahunakademik"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "X8.Master Setting Periode",
			"item": [
				{
					"name": "List Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNDg5NjY4fQ.rze0BLyiqiT5OLVIWnE_VVBHSHB8vkE34jsd84XRvBY",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/settingperiode/listsettperiode",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"settingperiode",
								"listsettperiode"
							]
						}
					},
					"response": []
				},
				{
					"name": "Tampil Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNDg5NjY4fQ.rze0BLyiqiT5OLVIWnE_VVBHSHB8vkE34jsd84XRvBY",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/settingperiode/showsettperiode?page=1&perpage=10&sort=asc",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"settingperiode",
								"showsettperiode"
							],
							"query": [
								{
									"key": "page",
									"value": "1"
								},
								{
									"key": "perpage",
									"value": "10"
								},
								{
									"key": "search",
									"value": "2023",
									"disabled": true
								},
								{
									"key": "sort",
									"value": "asc"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "Create Data",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NjUwfQ.QfqH5geTCcj6YhpO3LTqILTGJ-g4Vm3iHfGNkWn_L6A",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_periode_spp\" : 1,\r\n    \"nm_kelas\" : \"XII\",\r\n    \"biaya_spp\" : 350000,\r\n    \"keterangan\" : \"Biaya SPP Tahun Akademik 2022-2023 Kelas Dua Belas\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/settingperiode/insertsettperiode",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"settingperiode",
								"insertsettperiode"
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNDg5NjY4fQ.rze0BLyiqiT5OLVIWnE_VVBHSHB8vkE34jsd84XRvBY",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_periode_spp\" : 1,\r\n    \"nm_kelas\" : \"X\",\r\n    \"biaya_spp\" : 300000,\r\n    \"keterangan\" : \"Biaya SPP Tahun Akademik 2022-2023 Kelas Sepuluh\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/settingperiode/updatesettperiode/1",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"settingperiode",
								"updatesettperiode",
								"1"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNDg5NjY4fQ.rze0BLyiqiT5OLVIWnE_VVBHSHB8vkE34jsd84XRvBY",
								"type": "text"
							}
						],
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/settingperiode/deletesettperiode/1",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"settingperiode",
								"deletesettperiode",
								"1"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "8.Master Kelas Akademik",
			"item": [
				{
					"name": "List Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/akademik/listkelasakademik",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"akademik",
								"listkelasakademik"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "9.Transaksi UM SPP",
			"item": [
				{
					"name": "1. List Kd Group",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukspp/listgroupkategori",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukspp",
								"listgroupkategori"
							]
						}
					},
					"response": []
				},
				{
					"name": "2. List Kd Kategori",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukspp/listkategoriuang",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukspp",
								"listkategoriuang"
							]
						}
					},
					"response": []
				},
				{
					"name": "3. List Kelas",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukspp/listkelas",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukspp",
								"listkelas"
							]
						}
					},
					"response": []
				},
				{
					"name": "4. List Siswa",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tahun_akademik\" : \"2022/2023\",\r\n    \"nm_kelas\" : \"XIII\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://10.5.57.55:2022/api/v1/transaksi/uangmasukspp/listsiswa",
							"protocol": "http",
							"host": [
								"10",
								"5",
								"57",
								"55"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukspp",
								"listsiswa"
							]
						}
					},
					"response": []
				},
				{
					"name": "5. List Data",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tahun_akademik\" : \"2022/2023\",\r\n    \"nm_kelas\" : \"X\",\r\n    \"nis_siswa\" : \"5897\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukspp/listdata",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukspp",
								"listdata"
							]
						}
					},
					"response": []
				},
				{
					"name": "6. Create SPP Uang Masuk",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 3,\r\n    \"kd_kategori\" : 3,\r\n    \"tahun_akademik\" : \"2022/2023\",\r\n    \"nis_siswa\" : \"5664\",\r\n    \"nm_kelas\" : \"XII\",    \r\n    \"keterangan\" : \"SPP PEMBAYARAN ATAS NAMA DIANING PAKARTI\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukspp/createuangmasukspp/",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukspp",
								"createuangmasukspp",
								""
							]
						}
					},
					"response": []
				},
				{
					"name": "7. Edit SPP Uang Masuk",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tgl_bayar\" : \"02-07-2022\",\r\n    \"jml_bayar\" : 750000,\r\n    \"keterangan\" : \"test 777777\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukspp/updateuangmasukspp/:kd_trans_masuk/:kd_trans_masuk_detail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukspp",
								"updateuangmasukspp",
								":kd_trans_masuk",
								":kd_trans_masuk_detail"
							],
							"variable": [
								{
									"key": "kd_trans_masuk",
									"value": "6"
								},
								{
									"key": "kd_trans_masuk_detail",
									"value": "62"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "8. Delete All UM SPP",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukspp/deletealluangmasuk/:kd_trans_masuk",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukspp",
								"deletealluangmasuk",
								":kd_trans_masuk"
							],
							"variable": [
								{
									"key": "kd_trans_masuk",
									"value": "2"
								}
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "10.Master Siswa",
			"item": [
				{
					"name": "List Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/akademik/listsiswaakademik",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"akademik",
								"listsiswaakademik"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "11.Transaksi UM PPDB",
			"item": [
				{
					"name": "1. List Kd Group",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/listgroupkategori",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"listgroupkategori"
							]
						}
					},
					"response": []
				},
				{
					"name": "2. List Kd Kategori",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/listkategoriuang",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"listkategoriuang"
							]
						}
					},
					"response": []
				},
				{
					"name": "3. List Siswa",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/listsiswa",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"listsiswa"
							]
						}
					},
					"response": []
				},
				{
					"name": "4. List Data",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"nik\" : \"3172032402070005\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/listdata",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"listdata"
							]
						}
					},
					"response": []
				},
				{
					"name": "5. Create PPdb Uang Masuk",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 1,\r\n    \"kd_kategori\" : 1,\r\n    \"nik\" : \"3172031711060006\",\r\n    \"tgldaftar\" : \"28-06-2022\",\r\n    \"tahun_daftar\" : \"2022\",\r\n    \"tahun_akademik\" : \"2022-2023\", \r\n    \"keterangan\" : \"\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/createuangmasukppdb/",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"createuangmasukppdb",
								""
							]
						}
					},
					"response": []
				},
				{
					"name": "old6. Edit PPdb Uang Masuk",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tgl_bayar\" : \"02-07-2022\",\r\n    \"jml_bayar\" : 260000,\r\n    \"keterangan\" : \"hehehehe\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/updateuangmasukppdb/:kd_trans_masuk_ppdb/:kd_trans_masuk_detail_ppdb",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"updateuangmasukppdb",
								":kd_trans_masuk_ppdb",
								":kd_trans_masuk_detail_ppdb"
							],
							"variable": [
								{
									"key": "kd_trans_masuk_ppdb",
									"value": "1"
								},
								{
									"key": "kd_trans_masuk_detail_ppdb",
									"value": "3"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "6. Edit PPdb Uang Masuk Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tgl_bayar\" : \"02-09-2022\",\r\n    \"jml_bayar\" : 1000000,\r\n    \"kategori_biaya_ppdb\" : \"nyicil\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/updateuangmasukppdb/:kd_trans_masuk_ppdb/:kd_trans_masuk_detail_ppdb",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"updateuangmasukppdb",
								":kd_trans_masuk_ppdb",
								":kd_trans_masuk_detail_ppdb"
							],
							"variable": [
								{
									"key": "kd_trans_masuk_ppdb",
									"value": "1"
								},
								{
									"key": "kd_trans_masuk_detail_ppdb",
									"value": "3"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "7. Delete All UM PPdb",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/deletealluangmasuk/:kd_trans_masuk_ppdb",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"deletealluangmasuk",
								":kd_trans_masuk_ppdb"
							],
							"variable": [
								{
									"key": "kd_trans_masuk_ppdb",
									"value": "1"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "8. Add Uang Masuk Detail",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_masuk_ppdb\" : 1\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/createuangmasukdetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"createuangmasukdetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "9. Delete Uang Masuk Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_masuk_ppdb\" : 1,\r\n    \"kd_trans_masuk_detail_ppdb\" : 3\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/deleteuangmasukdetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"deleteuangmasukdetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "10. Edit Uang Masuk PPDB",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"keterangan\" : \"Edit Nominal Biaya PPDB\",\r\n    \"total_biaya\" : 5000000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/edituangmasuk/:kd_trans_masuk_ppdb",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"edituangmasuk",
								":kd_trans_masuk_ppdb"
							],
							"variable": [
								{
									"key": "kd_trans_masuk_ppdb",
									"value": "1"
								}
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "12.Transaksi UM Siswa",
			"item": [
				{
					"name": "1. List Kd Group",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuksiswa/listgroupkategori",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuksiswa",
								"listgroupkategori"
							]
						}
					},
					"response": []
				},
				{
					"name": "2. List Kd Kategori",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : \"2\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuksiswa/listkategoriuang",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuksiswa",
								"listkategoriuang"
							]
						}
					},
					"response": []
				},
				{
					"name": "3. List Add Siswa",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"nis\" : \"5897\",\r\n    \"kd_kategori\" : 17,\r\n    \"kd_group\" : 10\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuksiswa/listdataaddsiswa",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuksiswa",
								"listdataaddsiswa"
							]
						}
					},
					"response": []
				},
				{
					"name": "4. Create Siswa Uang Masuk",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 2,\r\n    \"kd_kategori\" : 2,\r\n    \"tahun_akademik\" : \"2022/2023\",\r\n    \"nis_siswa\" : \"5959\",\r\n    \"nm_kelas\" : \"X\",    \r\n    \"keterangan\" : \"\",\r\n    \"total_biaya\" : 0\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuksiswa/createuangmasuksiswa/",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuksiswa",
								"createuangmasuksiswa",
								""
							]
						}
					},
					"response": []
				},
				{
					"name": "5. Edit Uang Masuk Siswa",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 2,\r\n    \"kd_kategori\" : 2,\r\n    \"tahun_akademik\" : \"2022/2023\",\r\n    \"nis_siswa\" : \"5959\",\r\n    \"nm_kelas\" : \"X\",    \r\n    \"keterangan\" : \"TEST DOANG\",\r\n    \"total_biaya\" : 1600000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuksiswa/edituangmasuksiswa/:kd_trans_masuk_siswa",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuksiswa",
								"edituangmasuksiswa",
								":kd_trans_masuk_siswa"
							],
							"variable": [
								{
									"key": "kd_trans_masuk_siswa",
									"value": "4"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "6. Edit Uang Masuk Siswa Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tgl_bayar\" : \"02-07-2022\",\r\n    \"jml_bayar\" : 1000000,\r\n    \"keterangan\" : \"hehehehe\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuksiswa/updateuangmasuksiswadetail/:kd_trans_masuk_siswa/:kd_trans_masuk_detail_siswa",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuksiswa",
								"updateuangmasuksiswadetail",
								":kd_trans_masuk_siswa",
								":kd_trans_masuk_detail_siswa"
							],
							"variable": [
								{
									"key": "kd_trans_masuk_siswa",
									"value": "4"
								},
								{
									"key": "kd_trans_masuk_detail_siswa",
									"value": "7"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "7. Add Siswa Uang Masuk Detail",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_masuk_siswa\" : 4\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuksiswa/createuangmasuksiswadetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuksiswa",
								"createuangmasuksiswadetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "8. List Data (Dahsboard)",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tahun_akademik\" : \"2022/2023\",\r\n    \"nis_siswa\" : \"\",\r\n    \"nm_kelas\" : \"\"\r\n    \r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuksiswa/listdata",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuksiswa",
								"listdata"
							]
						}
					},
					"response": []
				},
				{
					"name": "9. Delete Uang Masuk Siswa Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_masuk_siswa\" : \"4\",\r\n    \"kd_trans_masuk_detail_siswa\" : \"5\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuksiswa/deleteuangmasuksiswadetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuksiswa",
								"deleteuangmasuksiswadetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "10. Delete All Um Siswa",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuksiswa/deletealluangmasuk/:kd_trans_masuk_siswa",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuksiswa",
								"deletealluangmasuk",
								":kd_trans_masuk_siswa"
							],
							"variable": [
								{
									"key": "kd_trans_masuk_siswa",
									"value": "1"
								}
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "13.Master Configurasi Spp PPdb",
			"item": [
				{
					"name": "Tampil Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterconfsppdb/showconfspppdb?page=1&perpage=10",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterconfsppdb",
								"showconfspppdb"
							],
							"query": [
								{
									"key": "page",
									"value": "1"
								},
								{
									"key": "perpage",
									"value": "10"
								},
								{
									"key": "search",
									"value": "na",
									"disabled": true
								},
								{
									"key": "sort",
									"value": "asc",
									"disabled": true
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 1,\r\n    \"kd_kategori\" : 1\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterconfsppdb/updateconfspppdb/2",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterconfsppdb",
								"updateconfspppdb",
								"2"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "14.Master Configurasi Biaya Kategori",
			"item": [
				{
					"name": "Tampil Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterbiayakategori/showbiayakategori?page=1&perpage=10",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterbiayakategori",
								"showbiayakategori"
							],
							"query": [
								{
									"key": "page",
									"value": "1"
								},
								{
									"key": "perpage",
									"value": "10"
								},
								{
									"key": "search",
									"value": "na",
									"disabled": true
								},
								{
									"key": "sort",
									"value": "asc",
									"disabled": true
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "List Data",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterbiayakategori/listbiayakategori",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterbiayakategori",
								"listbiayakategori"
							]
						}
					},
					"response": []
				},
				{
					"name": "Create Data",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_kategori\": 4,\r\n    \"jml_biaya\" : 220000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterbiayakategori/insertbiayakategori",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterbiayakategori",
								"insertbiayakategori"
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit Data",
					"request": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_kategori\": 3,\r\n    \"jml_biaya\" : 205000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterbiayakategori/updatebiayakategori/2",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterbiayakategori",
								"updatebiayakategori",
								"2"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Data",
					"request": {
						"method": "DELETE",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIn0.LBBjy6Ofdt6e-4RjgR9EiwDkxJRnHpU7XcO6DH7T2M0",
								"type": "text"
							}
						],
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/masterbiayakategori/deletebiayakategori/2",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"masterbiayakategori",
								"deletebiayakategori",
								"2"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "15.Transaksi UM Lain-Lain",
			"item": [
				{
					"name": "1. List Kd Group",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuklainlain/listgroupkategori",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuklainlain",
								"listgroupkategori"
							]
						}
					},
					"response": []
				},
				{
					"name": "2. List Kd Kategori",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : \"10\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuklainlain/listkategoriuang",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuklainlain",
								"listkategoriuang"
							]
						}
					},
					"response": []
				},
				{
					"name": "3. Create Uang Masuk Lain",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 10,\r\n    \"kd_kategori\" : 17,\r\n    \"no_document\" : \"TRF1232\",\r\n    \"tgl_document\" : \"02-01-2023\",  \r\n    \"keterangan\" : \"\",\r\n    \"total_biaya\" : 60000000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuklainlain/createuangmasuklain/",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuklainlain",
								"createuangmasuklain",
								""
							]
						}
					},
					"response": []
				},
				{
					"name": "4. Edit Uang Masuk Lain",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 10,\r\n    \"kd_kategori\" : 16,\r\n    \"no_document\" : \"nodok987263\",\r\n    \"tgl_document\" : \"12-12-2022\",\r\n    \"keterangan\" : \"Edit Uang Masuk\",\r\n    \"total_biaya\" : 5000000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuklainlain/edituangmasuklain/:kd_trans_masuk_lain",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuklainlain",
								"edituangmasuklain",
								":kd_trans_masuk_lain"
							],
							"variable": [
								{
									"key": "kd_trans_masuk_lain",
									"value": "2"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "5. Edit Uang Lain Siswa Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tgl_bayar\" : \"06-01-2023\",\r\n    \"jml_bayar\" : 60000000,\r\n    \"keterangan\" : \"hehehehe\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuklainlain/updateuangmasuklaindetail/:kd_trans_masuk_lain/:kd_trans_masuk_detail_lain",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuklainlain",
								"updateuangmasuklaindetail",
								":kd_trans_masuk_lain",
								":kd_trans_masuk_detail_lain"
							],
							"variable": [
								{
									"key": "kd_trans_masuk_lain",
									"value": "2"
								},
								{
									"key": "kd_trans_masuk_detail_lain",
									"value": "3"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "6. Add Lain Uang Masuk Detail",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_masuk_lain\" : 1\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuklainlain/createuangmasuklaindetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuklainlain",
								"createuangmasuklaindetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "7. List Data (Dahsboard)",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tgl_document1\" : \"01-01-2022\",\r\n    \"tgl_document2\" : \"31-12-2022\",\r\n    \"no_document\" : \"\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuklainlain/listdata",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuklainlain",
								"listdata"
							]
						}
					},
					"response": []
				},
				{
					"name": "8. Delete Uang Masuk Lain Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_masuk_lain\" : \"2\",\r\n    \"kd_trans_masuk_detail_lain\" : \"5\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuklainlain/deleteuangmasuklaindetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuklainlain",
								"deleteuangmasuklaindetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "9. Delete All UM Lain",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasuklainlain/deletealluangmasuk/:kd_trans_masuk_lain",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasuklainlain",
								"deletealluangmasuk",
								":kd_trans_masuk_lain"
							],
							"variable": [
								{
									"key": "kd_trans_masuk_lain",
									"value": "1"
								}
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "16.Transaksi Uang Keluar PRA",
			"item": [
				{
					"name": "1. List Kd Group",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpra/listgroupkategori",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpra",
								"listgroupkategori"
							]
						}
					},
					"response": []
				},
				{
					"name": "2. List Kd Kategori",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : \"9\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpra/listkategoriuang",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpra",
								"listkategoriuang"
							]
						}
					},
					"response": []
				},
				{
					"name": "3. Create Uang Keluar",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 9,\r\n    \"kd_kategori\" : 12,\r\n    \"no_document\" : \"nodok9887263\",\r\n    \"tgl_document\" : \"03-01-2023\",  \r\n    \"keterangan\" : \"\",\r\n    \"total_biaya\" : 2000000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpra/createuangkeluar",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpra",
								"createuangkeluar"
							]
						}
					},
					"response": []
				},
				{
					"name": "4. Edit Uang Keluar",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 9,\r\n    \"kd_kategori\" : 12,\r\n    \"kd_proses\" : \"PRA\",\r\n    \"no_document\" : \"nodok987263\",\r\n    \"tgl_document\" : \"12-12-2022\",\r\n    \"keterangan\" : \"Edit Uang Keluar\",\r\n    \"total_biaya\" : 1900000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpra/edituangkeluar/:kd_trans_keluar",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpra",
								"edituangkeluar",
								":kd_trans_keluar"
							],
							"variable": [
								{
									"key": "kd_trans_keluar",
									"value": "1"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "5. Edit Uang Keluar Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"jml_bayar\" : 2000000,\r\n    \"keterangan\" : \"hehehehe\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpra/updateuangkeluardetail/:kd_trans_keluar/:kd_trans_keluar_detail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpra",
								"updateuangkeluardetail",
								":kd_trans_keluar",
								":kd_trans_keluar_detail"
							],
							"variable": [
								{
									"key": "kd_trans_keluar",
									"value": "2"
								},
								{
									"key": "kd_trans_keluar_detail",
									"value": "3"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "6. Add Uang Keluar Detail",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_keluar\" : 1\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpra/createuangkeluardetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpra",
								"createuangkeluardetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "7. List Data (Dahsboard)",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tgl_document1\" : \"12-12-2022\",\r\n    \"tgl_document2\" : \"31-12-2022\",\r\n    \"no_document\" : \"\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpra/listdata",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpra",
								"listdata"
							]
						}
					},
					"response": []
				},
				{
					"name": "8. Delete Uang Keluar Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_keluar\" : \"1\",\r\n    \"kd_trans_keluar_detail\" : \"2\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpra/deleteuangkeluardetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpra",
								"deleteuangkeluardetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "9. Delete All Uang Keluar",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpra/deletealluangkeluar/:kd_trans_keluar",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpra",
								"deletealluangkeluar",
								":kd_trans_keluar"
							],
							"variable": [
								{
									"key": "kd_trans_keluar",
									"value": "1"
								}
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "17.Transaksi Uang Keluar PRA-ACT",
			"item": [
				{
					"name": "1. List Dokument PRA",
					"request": {
						"method": "GET",
						"header": [
							{
								"warning": "This is a duplicate header and will be overridden by the Authorization header generated by Postman.",
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpraact/listdokument",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpraact",
								"listdokument"
							]
						}
					},
					"response": []
				},
				{
					"name": "2. Post Uang Masuk",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpraact/listgroupkategori",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpraact",
								"listgroupkategori"
							]
						}
					},
					"response": []
				},
				{
					"name": "3. Create Uang Keluar",
					"request": {
						"method": "POST",
						"header": [
							{
								"warning": "This is a duplicate header and will be overridden by the Authorization header generated by Postman.",
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_keluar\" : 1,\r\n    \"no_document\" : \"nodok987263\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpraact/createuangkeluar",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpraact",
								"createuangkeluar"
							]
						}
					},
					"response": []
				},
				{
					"name": "4. Edit Uang Keluar",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"total_biaya\" : 5500000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpraact/edituangkeluar/:kd_trans_keluar",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpraact",
								"edituangkeluar",
								":kd_trans_keluar"
							],
							"variable": [
								{
									"key": "kd_trans_keluar",
									"value": "1"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "5. Edit Uang Keluar Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_post_uang_masuk\" : 10,\r\n    \"tgl_bayar\" : \"05-01-2023\",\r\n    \"jml_bayar\" : 1900000,\r\n    \"keterangan\" : \"Ambil dari Uang Masuk Eksternal\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpraact/updateuangkeluardetail/:kd_trans_keluar/:kd_trans_keluar_detail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpraact",
								"updateuangkeluardetail",
								":kd_trans_keluar",
								":kd_trans_keluar_detail"
							],
							"variable": [
								{
									"key": "kd_trans_keluar",
									"value": "1"
								},
								{
									"key": "kd_trans_keluar_detail",
									"value": "1"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "6. Add Uang Keluar Detail",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_keluar\" : 2\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpraact/createuangkeluardetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpraact",
								"createuangkeluardetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "7. List Data (Dahsboard)",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tgl_document1\" : \"12-12-2022\",\r\n    \"tgl_document2\" : \"31-12-2023\",\r\n    \"no_document\" : \"\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpraact/listdata",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpraact",
								"listdata"
							]
						}
					},
					"response": []
				},
				{
					"name": "8. Delete Uang Keluar Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_keluar\" : \"1\",\r\n    \"kd_trans_keluar_detail\" : \"2\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpraact/deleteuangkeluardetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpraact",
								"deleteuangkeluardetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "9. Delete All Uang Keluar",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpraact/deletealluangkeluar/:kd_trans_keluar",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpraact",
								"deletealluangkeluar",
								":kd_trans_keluar"
							],
							"variable": [
								{
									"key": "kd_trans_keluar",
									"value": "1"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "10. Post Uang Masuk",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluarpraact/postuangmasuk",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluarpraact",
								"postuangmasuk"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "18.Transaksi Uang Keluar ACT",
			"item": [
				{
					"name": "1. List Kd Group",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluaract/listgroupkategori",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluaract",
								"listgroupkategori"
							]
						}
					},
					"response": []
				},
				{
					"name": "2. List Kd Kategori",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : \"7\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluaract/listkategoriuang",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluaract",
								"listkategoriuang"
							]
						}
					},
					"response": []
				},
				{
					"name": "3. Create Uang Keluar",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 8,\r\n    \"kd_kategori\" : 8,\r\n    \"no_document\" : \"INV-PLN-001\",\r\n    \"tgl_document\" : \"20-01-2023\",  \r\n    \"keterangan\" : \"Bayar PLN Listrik\",\r\n    \"total_biaya\" : 4500000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluaract/createuangkeluar",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluaract",
								"createuangkeluar"
							]
						}
					},
					"response": []
				},
				{
					"name": "4. Edit Uang Keluar",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_group\" : 7,\r\n    \"kd_kategori\" : 7,\r\n    \"no_document\" : \"gaji982983z\",\r\n    \"tgl_document\" : \"06-01-2023\",\r\n    \"keterangan\" : \"Edit Uang Keluar\",\r\n    \"total_biaya\" :1200000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluaract/edituangkeluar/:kd_trans_keluar",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluaract",
								"edituangkeluar",
								":kd_trans_keluar"
							],
							"variable": [
								{
									"key": "kd_trans_keluar",
									"value": "1"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "5. Edit Uang Keluar Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_post_uang_masuk\" : 10,\r\n    \"tgl_bayar\" : \"05-01-2023\",\r\n    \"jml_bayar\" : 8700000,\r\n    \"keterangan\" : \"Bayar Listrik Periode Desember 2022\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluaract/updateuangkeluardetail/:kd_trans_keluar/:kd_trans_keluar_detail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluaract",
								"updateuangkeluardetail",
								":kd_trans_keluar",
								":kd_trans_keluar_detail"
							],
							"variable": [
								{
									"key": "kd_trans_keluar",
									"value": "2"
								},
								{
									"key": "kd_trans_keluar_detail",
									"value": "4"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "6. Add Uang Keluar Detail",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_keluar\" : 2\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluaract/createuangkeluardetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluaract",
								"createuangkeluardetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "7. List Data (Dahsboard)",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tgl_document1\" : \"12-01-2022\",\r\n    \"tgl_document2\" : \"31-12-2023\",\r\n    \"no_document\" : \"\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluaract/listdata",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluaract",
								"listdata"
							]
						}
					},
					"response": []
				},
				{
					"name": "8. Delete Uang Keluar Detail",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kd_trans_keluar\" : \"1\",\r\n    \"kd_trans_keluar_detail\" : \"3\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluaract/deleteuangkeluardetail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluaract",
								"deleteuangkeluardetail"
							]
						}
					},
					"response": []
				},
				{
					"name": "9. Delete All Uang Keluar",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangkeluaract/deletealluangkeluar/:kd_trans_keluar",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangkeluaract",
								"deletealluangkeluar",
								":kd_trans_keluar"
							],
							"variable": [
								{
									"key": "kd_trans_keluar",
									"value": "1"
								}
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "19.Report Uang Masuk",
			"item": [
				{
					"name": "Report SPP",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tahun_akademik\" : \"2022/2023\",\r\n    \"nis_siswa\" : \"5897\",\r\n    \"periode_bayar1\" : \"01-2022\" ,\r\n    \"periode_bayar2\" : \"12-2023\" ,\r\n    \"tgl_bayar1\" : \"01-01-2022\",\r\n    \"tgl_bayar2\" : \"31-12-2022\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/report/uangmasuk/reportspp",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"report",
								"uangmasuk",
								"reportspp"
							]
						}
					},
					"response": []
				},
				{
					"name": "Report PPDB",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tahun_akademik\" : \"2022-2023\",\r\n    \"nik\" : \"\",\r\n    \"tgldaftar1\" : \"01-01-2022\",\r\n    \"tgldaftar2\" : \"01-01-2024\",\r\n    \"tgl_bayar1\" : \"01-01-2022\",\r\n    \"tgl_bayar2\" : \"31-12-2022\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/report/uangmasuk/reportppdb",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"report",
								"uangmasuk",
								"reportppdb"
							]
						}
					},
					"response": []
				},
				{
					"name": "Report UM Siswa",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tahun_akademik\" : \"2022/2023\",\r\n    \"nis_siswa\" : \"\",\r\n    \"nm_kelas\" : \"\" ,\r\n    \"tgl_bayar1\" : \"01-01-2022\",\r\n    \"tgl_bayar2\" : \"31-12-2023\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/report/uangmasuk/reportumsiswa",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"report",
								"uangmasuk",
								"reportumsiswa"
							]
						}
					},
					"response": []
				},
				{
					"name": "Report UM Lain",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"no_document\" : \"nodok987263\",\r\n    \"tgl_document1\" : \"01-01-2022\",\r\n    \"tgl_document2\" : \"31-12-2023\",\r\n    \"tgl_bayar1\" : \"01-01-2022\",\r\n    \"tgl_bayar2\" : \"31-12-2023\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/report/uangmasuk/reportumlain",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"report",
								"uangmasuk",
								"reportumlain"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "20.Report Uang Keluar",
			"item": [
				{
					"name": "Report PRA",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"no_document\" : \"nodok98726\",\r\n    \"tgl_document1\" : \"01-01-2022\",\r\n    \"tgl_document2\" : \"31-12-2023\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/report/uangkeluar/reportpra",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"report",
								"uangkeluar",
								"reportpra"
							]
						}
					},
					"response": []
				},
				{
					"name": "Report PRA ACT",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"no_document\" : \"\",\r\n    \"tgl_document1\" : \"01-01-2022\",\r\n    \"tgl_document2\" : \"31-12-2023\",\r\n    \"tgl_bayar1\" : \"\",\r\n    \"tgl_bayar2\" : \"\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/report/uangkeluar/reportpraact",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"report",
								"uangkeluar",
								"reportpraact"
							]
						}
					},
					"response": []
				},
				{
					"name": "Report ACT",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"no_document\" : \"\",\r\n    \"tgl_document1\" : \"01-01-2022\",\r\n    \"tgl_document2\" : \"31-12-2023\",\r\n    \"tgl_bayar1\" : \"01-01-2022\",\r\n    \"tgl_bayar2\" : \"31-12-2023\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/report/uangkeluar/reportact",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"report",
								"uangkeluar",
								"reportact"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "21. Import all Data PPDB",
			"item": [
				{
					"name": "1. List Data PPDB",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tahun_daftar\" : \"2022\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/datalistppdb",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"datalistppdb"
							]
						}
					},
					"response": []
				},
				{
					"name": "2. Import Data",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tahun_daftar\" : \"2022\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukppdb/importallppdb",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukppdb",
								"importallppdb"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "New Folder",
			"item": [
				{
					"name": "3. List Kelas",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukspp/listkelas",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukspp",
								"listkelas"
							]
						}
					},
					"response": []
				},
				{
					"name": "5. List Data",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZF91c2VyIjoxLCJVc2VybmFtZSI6InRlZ3VoIiwiZXhwIjoxNjYyNzU2NTA1fQ.LWQuFuoT-accZZK0iDgoCy2Jx8NI_URnFsDNofk1dFM",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tahun_akademik\" : \"2022-2023\",\r\n    \"nm_kelas\" : \"XI\",\r\n    \"nis_siswa\" : \"5762\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukspp/listdata",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukspp",
								"listdata"
							]
						}
					},
					"response": []
				},
				{
					"name": "7. Edit SPP Uang Masuk",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"tgl_bayar\" : \"02-07-2022\",\r\n    \"jml_bayar\" : 100000,\r\n    \"keterangan\" : \"hehehehe\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:2022/api/v1/transaksi/uangmasukspp/updateuangmasukspp/:kd_trans_masuk/:kd_trans_masuk_detail",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "2022",
							"path": [
								"api",
								"v1",
								"transaksi",
								"uangmasukspp",
								"updateuangmasukspp",
								":kd_trans_masuk",
								":kd_trans_masuk_detail"
							],
							"variable": [
								{
									"key": "kd_trans_masuk",
									"value": "1"
								},
								{
									"key": "kd_trans_masuk_detail",
									"value": "6"
								}
							]
						}
					},
					"response": []
				}
			]
		}
	],
	"auth": {
		"type": "bearer",
		"bearer": [
			{
				"key": "token",
				"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGdWxsX25hbWUiOiJUZWd1aCBTdWdpb25vIiwiSWRfdXNlciI6MSwiVXNlcm5hbWUiOiJ0ZWd1aCIsImV4cCI6MTY3NjA4NTA2OX0.nCy_9m-45AWyzb5hug9wojo5ans7rSm-MB1CzicQoiw",
				"type": "string"
			}
		]
	},
	"event": [
		{
			"listen": "prerequest",
			"script": {
				"type": "text/javascript",
				"exec": [
					""
				]
			}
		},
		{
			"listen": "test",
			"script": {
				"type": "text/javascript",
				"exec": [
					""
				]
			}
		}
	]
}