/*
Navicat MySQL Data Transfer

Source Server         : localhost3306-MBA
Source Server Version : 50531
Source Host           : localhost:3306
Source Database       : cuangshu_db_bendahara

Target Server Type    : MYSQL
Target Server Version : 50531
File Encoding         : 65001

Date: 2022-09-04 12:09:53
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `tbl_conf_periode_spps`
-- ----------------------------
DROP TABLE IF EXISTS `tbl_conf_periode_spps`;
CREATE TABLE `tbl_conf_periode_spps` (
  `kd_periode_spp` int(11) DEFAULT NULL,
  `seqno` int(11) DEFAULT NULL,
  `kd_bulan` char(2) DEFAULT NULL,
  `tahun` int(11) DEFAULT NULL,
  `nm_sett` varchar(200) DEFAULT NULL,
  `flag_aktif` int(1) DEFAULT NULL,
  `created_on` datetime DEFAULT NULL,
  `created_by` varchar(100) DEFAULT NULL,
  `edited_on` datetime DEFAULT NULL,
  `edited_by` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tbl_conf_periode_spps
-- ----------------------------

-- ----------------------------
-- Table structure for `tbl_group_kategoris`
-- ----------------------------
DROP TABLE IF EXISTS `tbl_group_kategoris`;
CREATE TABLE `tbl_group_kategoris` (
  `kd_jenis` int(11) DEFAULT NULL,
  `kd_group` int(11) NOT NULL AUTO_INCREMENT,
  `nm_group` varchar(100) DEFAULT NULL,
  `nm_header` varchar(500) DEFAULT NULL,
  `nm_detail` varchar(500) DEFAULT NULL,
  `flag_aktif` int(1) DEFAULT '0',
  `created_on` datetime DEFAULT NULL,
  `created_by` varchar(100) DEFAULT NULL,
  `edited_on` datetime DEFAULT NULL,
  `edited_by` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`kd_group`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tbl_group_kategoris
-- ----------------------------
INSERT INTO `tbl_group_kategoris` VALUES ('1', '1', 'Biaya Pembayaran PPDB', 'PPDB,Daftar Ulang,SPP,Semester', 'Biaya Seragam,Biaya Uang Gedung,Biaya Formulir Pendaftaran', '0', '2022-08-27 20:08:58', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('1', '2', 'Biaya Pembayaran Daftar Ulang', null, null, '0', '2022-08-27 20:09:10', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('1', '3', 'Biaya Pembayaran SPP', null, null, '0', '2022-08-27 20:09:33', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('1', '4', 'Biaya Pembayaran Semester', null, null, '0', '2022-08-27 20:09:45', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('1', '5', 'Biaya Pembayaran Lain-Lain', null, null, '0', '2022-08-27 20:10:26', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('2', '6', 'Biaya Gaji Pegawai', 'Gaji Guru,Gaji Staff', null, '0', '2022-08-27 20:10:38', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('2', '7', 'Biaya Kegiatan Siswa', null, null, '0', '2022-08-27 20:10:48', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('2', '8', 'Biaya Rutin Gedung Sekolah', null, null, '0', '2022-08-27 20:11:09', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('2', '9', 'Biaya Perlengkapan Siswa', null, null, '0', '2022-08-27 20:11:20', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('1', '10', 'Biaya Makan Anak Buaya', '', null, '9', '2022-08-27 20:12:00', 'teguh', '2022-08-28 01:44:52', 'teguh');

-- ----------------------------
-- Table structure for `tbl_jenis_trans`
-- ----------------------------
DROP TABLE IF EXISTS `tbl_jenis_trans`;
CREATE TABLE `tbl_jenis_trans` (
  `kd_jenis` int(11) NOT NULL AUTO_INCREMENT,
  `proses_uang` varchar(100) DEFAULT NULL,
  `flag_aktif` int(1) DEFAULT '0',
  `created_on` datetime DEFAULT NULL,
  `created_by` varchar(100) DEFAULT NULL,
  `edited_on` datetime DEFAULT NULL,
  `edited_by` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`kd_jenis`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tbl_jenis_trans
-- ----------------------------
INSERT INTO `tbl_jenis_trans` VALUES ('1', 'Uang Masuk', '0', '2022-08-27 20:04:15', 'teguh', null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('2', 'Uang Keluar', '0', '2022-08-27 20:04:39', 'teguh', null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('3', 'Uang Panas', '9', '2022-08-27 20:05:01', 'teguh', '2022-08-27 21:24:42', 'teguh');

-- ----------------------------
-- Table structure for `tbl_kategori_uangs`
-- ----------------------------
DROP TABLE IF EXISTS `tbl_kategori_uangs`;
CREATE TABLE `tbl_kategori_uangs` (
  `kd_group` int(11) DEFAULT NULL,
  `kd_kategori` int(11) NOT NULL AUTO_INCREMENT,
  `nm_kategori` varchar(200) DEFAULT NULL,
  `flag_aktif` int(1) DEFAULT NULL,
  `created_on` datetime DEFAULT NULL,
  `created_by` varchar(100) DEFAULT NULL,
  `edited_on` datetime DEFAULT NULL,
  `edited_by` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`kd_kategori`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tbl_kategori_uangs
-- ----------------------------
INSERT INTO `tbl_kategori_uangs` VALUES ('1', '1', 'PPDB', '0', '2022-08-27 23:33:46', 'teguh', null, null);
INSERT INTO `tbl_kategori_uangs` VALUES ('1', '2', 'Daftar Ulang', '0', '2022-08-27 23:34:14', 'teguh', '2022-08-28 01:36:45', 'teguh');
INSERT INTO `tbl_kategori_uangs` VALUES ('1', '3', 'SPP', '0', '2022-08-27 23:34:28', 'teguh', null, null);
INSERT INTO `tbl_kategori_uangs` VALUES ('1', '4', 'Semester', '0', '2022-08-27 23:34:39', 'teguh', null, null);
INSERT INTO `tbl_kategori_uangs` VALUES ('6', '5', 'Gaji Guru', '0', '2022-08-27 23:36:03', 'teguh', null, null);
INSERT INTO `tbl_kategori_uangs` VALUES ('6', '6', 'Gaji Staff', '0', '2022-08-27 23:36:27', 'teguh', '2022-08-28 01:36:16', 'teguh');
INSERT INTO `tbl_kategori_uangs` VALUES ('10', '7', 'Beli Anak Ayam', '9', '2022-08-28 01:38:51', 'teguh', '2022-08-28 02:51:25', 'teguh');
INSERT INTO `tbl_kategori_uangs` VALUES ('10', '8', 'Beli Anak Kambing', '9', '2022-08-28 01:39:15', 'teguh', '2022-08-28 01:41:51', 'teguh');

-- ----------------------------
-- Table structure for `tbl_sub_kategori_uangs`
-- ----------------------------
DROP TABLE IF EXISTS `tbl_sub_kategori_uangs`;
CREATE TABLE `tbl_sub_kategori_uangs` (
  `kd_kategori` int(11) DEFAULT NULL,
  `kd_sub_kategori` int(11) NOT NULL AUTO_INCREMENT,
  `nm_sub_kategori` varchar(200) DEFAULT NULL,
  `flag_aktif` int(11) DEFAULT NULL,
  `created_on` datetime DEFAULT NULL,
  `created_by` varchar(100) DEFAULT NULL,
  `edited_on` datetime DEFAULT NULL,
  `edited_by` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`kd_sub_kategori`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tbl_sub_kategori_uangs
-- ----------------------------
INSERT INTO `tbl_sub_kategori_uangs` VALUES ('1', '1', 'Biaya Seragam', '0', '2022-08-28 02:32:23', 'teguh', null, null);
INSERT INTO `tbl_sub_kategori_uangs` VALUES ('1', '2', 'Biaya Uang Gedung', '0', '2022-08-28 02:33:46', 'teguh', null, null);
INSERT INTO `tbl_sub_kategori_uangs` VALUES ('1', '3', 'Biaya Formulir Pendaftaran', '0', '2022-08-28 02:34:07', 'teguh', '2022-08-28 02:48:58', 'teguh');
INSERT INTO `tbl_sub_kategori_uangs` VALUES ('1', '4', 'Biaya Formulir Pendaftaran xxx', '9', '2022-08-28 02:49:43', 'teguh', '2022-08-28 02:50:02', 'teguh');

-- ----------------------------
-- Table structure for `tbl_tahun_akademik`
-- ----------------------------
DROP TABLE IF EXISTS `tbl_tahun_akademik`;
CREATE TABLE `tbl_tahun_akademik` (
  `id_tahun` int(11) NOT NULL AUTO_INCREMENT,
  `tahun_akademik` varchar(20) DEFAULT NULL,
  `status` varchar(15) DEFAULT NULL,
  `created_on` datetime DEFAULT NULL,
  `created_by` varchar(50) DEFAULT NULL,
  `edited_on` datetime DEFAULT NULL,
  `edited_by` varchar(50) DEFAULT NULL,
  `flag_tahun` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id_tahun`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tbl_tahun_akademik
-- ----------------------------
INSERT INTO `tbl_tahun_akademik` VALUES ('1', '2019/2020', 'Tidak Aktif', '2020-02-14 09:50:02', 'cuang', null, null, '0');
INSERT INTO `tbl_tahun_akademik` VALUES ('2', '2020/2021', 'Tidak Aktif', '2020-03-31 12:07:42', 'farhan', null, null, '0');
INSERT INTO `tbl_tahun_akademik` VALUES ('3', '2021/2022', 'Aktif', '2021-06-17 09:37:53', 'cuang', '2021-07-07 08:09:37', 'jejen', '0');
INSERT INTO `tbl_tahun_akademik` VALUES ('4', '2022-2023', 'Tidak Aktif', '2022-07-07 11:51:51', 'jejen', null, null, '0');

-- ----------------------------
-- Table structure for `tbl_users`
-- ----------------------------
DROP TABLE IF EXISTS `tbl_users`;
CREATE TABLE `tbl_users` (
  `id_user` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(200) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL,
  `created_on` datetime DEFAULT NULL,
  `edited_on` datetime DEFAULT NULL,
  PRIMARY KEY (`id_user`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tbl_users
-- ----------------------------
INSERT INTO `tbl_users` VALUES ('1', 'teguh', '$2a$04$hqeGToL4EuyL.sOu3dV4I.7GsCpN624ckPL.qrKUDh5x4v3F8ulli', '2022-08-06 04:31:37', null);
