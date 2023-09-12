-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 12, 2023 at 11:02 AM
-- Server version: 8.0.30
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `setik`
--

-- --------------------------------------------------------

--
-- Table structure for table `masyarakat`
--

CREATE TABLE `masyarakat` (
  `idm` int NOT NULL,
  `nik` bigint DEFAULT NULL,
  `nama` varchar(50) DEFAULT NULL,
  `no_hp` varchar(30) DEFAULT NULL,
  `gender` enum('laki-laki','perempuan') DEFAULT 'laki-laki',
  `tempat_lahir` varchar(50) NOT NULL,
  `alamat` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `birthday` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `masyarakat`
--

INSERT INTO `masyarakat` (`idm`, `nik`, `nama`, `no_hp`, `gender`, `tempat_lahir`, `alamat`, `created_at`, `update_at`, `birthday`) VALUES
(81, 3309018922540001, 'Jidane Adi Ramadhzan', '082134147290', 'laki-laki', 'Boyolali', 'Tegalarum Rt09/Rw01, Sumbung, Cepogo, Boyolali', '2023-08-09 21:09:14', '2023-09-11 07:37:15', '2001-06-11'),
(86, 3309018922540003, 'Mierta Ivani Choirunnisa', '082134147291', 'laki-laki', 'Boyolali', 'Tegalarum Rt09 Rw01, Sumbung, Cepogo, Boyolali', '2023-08-09 21:25:59', '2023-08-13 04:33:50', '1999-10-22'),
(87, 3309018932540002, 'Tirta Aura Ramazhan', '082134147290', 'laki-laki', 'Boyolali', 'Tegalarum Rt09 Rw01, Sumbung, Cepogo, Boyolali', '2023-08-09 21:28:22', '2023-08-13 05:08:39', '2003-12-24'),
(101, 3309019902833892, 'User ujian', '089338753889', 'laki-laki', 'Boyolali', 'Boyolali', '2023-08-18 22:26:58', '2023-08-18 22:26:58', '1997-02-05'),
(102, 3309028992003902, 'Ucob skuy', '089289392012', 'laki-laki', 'Boyolali', 'Boyolali', '2023-09-07 21:10:59', '2023-09-07 21:10:59', '2023-09-06'),
(106, 3309012345678901, 'Patar ', '089022010990', 'laki-laki', 'Sulawesi Tengah', 'Gedongan, Sinduadi, Mlati, Sleman', '2023-09-11 00:38:46', '2023-09-11 00:38:46', '2001-02-28'),
(107, 3309030120300002, 'Andi Sarwoedi', '082111089000', 'laki-laki', 'Magetan', 'Gedongan, Sinduadi, Mlati, Sleman, Daerah Istimewa Yogyakarta', '2023-09-11 18:42:47', '2023-09-11 18:42:47', '2001-05-02');

-- --------------------------------------------------------

--
-- Table structure for table `pengantar_kelahiran`
--

CREATE TABLE `pengantar_kelahiran` (
  `id` int NOT NULL,
  `id_surat` int NOT NULL,
  `ket_kelahiran_rs` varchar(50) NOT NULL,
  `dokumen_lain` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `pengantar_kematian`
--

CREATE TABLE `pengantar_kematian` (
  `id` int NOT NULL,
  `id_surat` int NOT NULL,
  `ket_kematian_dokter` varchar(50) NOT NULL,
  `sptjm` varchar(50) NOT NULL,
  `ktp_pelapor` varchar(50) NOT NULL,
  `ktp_saksi` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `pengantar_kk`
--

CREATE TABLE `pengantar_kk` (
  `id` int NOT NULL,
  `id_surat` int NOT NULL,
  `surat_pindah` varchar(50) NOT NULL,
  `fc_buku_nikah` varchar(50) NOT NULL,
  `fc_ktp_lama` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `pengantar_ktp`
--

CREATE TABLE `pengantar_ktp` (
  `id` int NOT NULL,
  `id_surat` int NOT NULL,
  `kk_lama` varchar(50) NOT NULL,
  `akte` varchar(50) NOT NULL,
  `ijazah` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `pengantar_tdkmampu`
--

CREATE TABLE `pengantar_tdkmampu` (
  `id` int NOT NULL,
  `id_surat` int NOT NULL,
  `kk` int NOT NULL,
  `ktp` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `surat`
--

CREATE TABLE `surat` (
  `id` int NOT NULL,
  `id_masyarakat` int NOT NULL DEFAULT '0',
  `jns_surat` enum('ktp','kk','kematian','kelahiran','tidak mampu') NOT NULL DEFAULT 'ktp',
  `status` enum('terverifikasi','diproses','ditolak','diterbitkan','diambil') NOT NULL DEFAULT 'diproses',
  `keterangan` text NOT NULL,
  `dokumen_syarat` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` bigint NOT NULL,
  `email` varchar(50) DEFAULT NULL,
  `password` text,
  `konf_pass` text,
  `role` enum('admin','masyarakat') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'masyarakat'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `email`, `password`, `konf_pass`, `role`) VALUES
(3309012345678901, 'patar@gmail.com', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', 'masyarakat'),
(3309012345678911, 'admin@gmail.com', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', 'admin'),
(3309018922540001, 'jidanear@gmail.com', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', 'masyarakat'),
(3309018922540003, 'mierta@gmail.com', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', 'masyarakat'),
(3309018932540002, 'tirtaar@gmail.com', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', 'masyarakat'),
(3309019902833892, 'user@yahoo.co.id', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', 'masyarakat'),
(3309028992003902, 'ucob@gmail.co.id', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', 'masyarakat'),
(3309030120300002, 'update@gmail.com', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', '8618ea1de89c78504a61b41491309f60e4f4e117b209255e26d835608bc96bcd', 'masyarakat');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `masyarakat`
--
ALTER TABLE `masyarakat`
  ADD PRIMARY KEY (`idm`),
  ADD KEY `NIK` (`nik`);

--
-- Indexes for table `pengantar_kelahiran`
--
ALTER TABLE `pengantar_kelahiran`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ID_KELAHIRAN` (`id_surat`);

--
-- Indexes for table `pengantar_kematian`
--
ALTER TABLE `pengantar_kematian`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ID_KEMATIAN` (`id_surat`);

--
-- Indexes for table `pengantar_kk`
--
ALTER TABLE `pengantar_kk`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ID_KK` (`id_surat`);

--
-- Indexes for table `pengantar_ktp`
--
ALTER TABLE `pengantar_ktp`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ID_KTP` (`id_surat`);

--
-- Indexes for table `pengantar_tdkmampu`
--
ALTER TABLE `pengantar_tdkmampu`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ID_TDKMAMPU` (`id_surat`);

--
-- Indexes for table `surat`
--
ALTER TABLE `surat`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_masyarakat` (`id_masyarakat`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `masyarakat`
--
ALTER TABLE `masyarakat`
  MODIFY `idm` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=108;

--
-- AUTO_INCREMENT for table `pengantar_kelahiran`
--
ALTER TABLE `pengantar_kelahiran`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `pengantar_kematian`
--
ALTER TABLE `pengantar_kematian`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `pengantar_kk`
--
ALTER TABLE `pengantar_kk`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `pengantar_ktp`
--
ALTER TABLE `pengantar_ktp`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `pengantar_tdkmampu`
--
ALTER TABLE `pengantar_tdkmampu`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `surat`
--
ALTER TABLE `surat`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `masyarakat`
--
ALTER TABLE `masyarakat`
  ADD CONSTRAINT `NIK` FOREIGN KEY (`nik`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `pengantar_kelahiran`
--
ALTER TABLE `pengantar_kelahiran`
  ADD CONSTRAINT `ID_KELAHIRAN` FOREIGN KEY (`id_surat`) REFERENCES `surat` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `pengantar_kematian`
--
ALTER TABLE `pengantar_kematian`
  ADD CONSTRAINT `ID_KEMATIAN` FOREIGN KEY (`id_surat`) REFERENCES `surat` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `pengantar_kk`
--
ALTER TABLE `pengantar_kk`
  ADD CONSTRAINT `ID_KK` FOREIGN KEY (`id_surat`) REFERENCES `surat` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `pengantar_ktp`
--
ALTER TABLE `pengantar_ktp`
  ADD CONSTRAINT `ID_KTP` FOREIGN KEY (`id_surat`) REFERENCES `surat` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `pengantar_tdkmampu`
--
ALTER TABLE `pengantar_tdkmampu`
  ADD CONSTRAINT `ID_TDKMAMPU` FOREIGN KEY (`id_surat`) REFERENCES `surat` (`id`);

--
-- Constraints for table `surat`
--
ALTER TABLE `surat`
  ADD CONSTRAINT `id_masyarakat` FOREIGN KEY (`id_masyarakat`) REFERENCES `masyarakat` (`idm`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
