/*
Navicat MySQL Data Transfer

Source Server         : localhost3306-MBA
Source Server Version : 50531
Source Host           : localhost:3306
Source Database       : cuangshu_db_bendahara

Target Server Type    : MYSQL
Target Server Version : 50531
File Encoding         : 65001

Date: 2022-08-25 14:04:04
*/

SET FOREIGN_KEY_CHECKS=0;

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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tbl_group_kategoris
-- ----------------------------
INSERT INTO `tbl_group_kategoris` VALUES ('1', '1', 'Biaya Pembayaran PPDB', null, null, '0', '2022-08-06 05:03:00', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('1', '2', 'Biaya Pembayaran Daftar Ulang', 'Semester,Eta Terangkanlah', null, '0', '2022-08-06 05:03:21', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('2', '3', 'Biaya Gaji Pegawai', null, null, '0', '2022-08-06 05:03:36', 'teguh', null, null);
INSERT INTO `tbl_group_kategoris` VALUES ('2', '4', 'Biaya Kegiatan Siswa', null, null, '0', '2022-08-06 05:03:48', 'teguh', '2022-08-06 05:16:23', 'teguh');
INSERT INTO `tbl_group_kategoris` VALUES ('2', '5', 'aaaaaaaaaaaaaa', null, null, '9', '2022-08-06 05:16:42', 'teguh', '2022-08-06 05:20:40', 'teguh');

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
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tbl_jenis_trans
-- ----------------------------
INSERT INTO `tbl_jenis_trans` VALUES ('1', 'Uang Masuk', '0', '2022-08-06 04:59:17', 'teguh', null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('2', 'Uang Keluar', '0', '2022-08-06 04:59:27', 'teguh', '2022-08-06 05:00:11', 'teguh');
INSERT INTO `tbl_jenis_trans` VALUES ('3', 'Uang Parkir', '0', '2022-08-06 04:59:31', 'teguh', '2022-08-06 05:00:36', 'teguh');
INSERT INTO `tbl_jenis_trans` VALUES ('4', 'Uang Suap', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('5', 'Uang Dinas', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('6', 'Uang Makan', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('7', 'Uang Jalan', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('8', 'Uang Gorengan', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('9', 'Uang Nasi Goreng', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('10', 'Uang Mie Goreng', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('11', 'Uang Jalan-Jalan', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('12', 'Uang Belanja', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('13', 'Uang Beli Mainan', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('14', 'Uang Beli Bubur', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('15', 'Uang Beli Ayam', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('16', 'Uang Beli Bebek', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('17', 'Uang Beli Burung', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('18', 'Uang Sodaqoh', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('19', 'Uang Sumbangan', '9', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('20', 'Uang Beli Aqua', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('21', 'Uang Jasa Parkir', '0', null, null, null, null);
INSERT INTO `tbl_jenis_trans` VALUES ('22', 'Uang Mainan', '0', null, null, null, null);

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
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tbl_kategori_uangs
-- ----------------------------
INSERT INTO `tbl_kategori_uangs` VALUES ('1', '1', 'PPDB', '0', '2022-08-07 22:30:20', 'teguh', '2022-08-07 22:47:49', 'teguh');
INSERT INTO `tbl_kategori_uangs` VALUES ('1', '2', 'Daftar Ulang', '0', '2022-08-07 22:32:00', 'teguh', null, null);
INSERT INTO `tbl_kategori_uangs` VALUES ('2', '4', 'Semester', '0', '2022-08-09 00:52:14', 'teguh', null, null);
INSERT INTO `tbl_kategori_uangs` VALUES ('1', '5', 'Semester', '0', '2022-08-09 01:23:43', 'teguh', null, null);
INSERT INTO `tbl_kategori_uangs` VALUES ('3', '6', 'Semester', '0', '2022-08-09 01:24:14', 'teguh', null, null);
INSERT INTO `tbl_kategori_uangs` VALUES ('1', '7', 'Semesterx', '0', '2022-08-09 01:39:27', 'teguh', null, null);
INSERT INTO `tbl_kategori_uangs` VALUES ('2', '17', 'Eta Terangkanlah', '0', '2022-08-13 00:13:19', 'teguh', null, null);

-- ----------------------------
-- Table structure for `tbl_sub_kategori_uangs`
-- ----------------------------
DROP TABLE IF EXISTS `tbl_sub_kategori_uangs`;
CREATE TABLE `tbl_sub_kategori_uangs` (
  `kd_kategori` int(11) DEFAULT NULL,
  `kd_sub_kategori` int(11) NOT NULL AUTO_INCREMENT,
  `nm_kategori` varchar(200) DEFAULT NULL,
  `flag_aktif` int(11) DEFAULT NULL,
  `created_on` datetime DEFAULT NULL,
  `created_by` varchar(100) DEFAULT NULL,
  `edited_on` datetime DEFAULT NULL,
  `edited_by` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`kd_sub_kategori`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tbl_sub_kategori_uangs
-- ----------------------------

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
