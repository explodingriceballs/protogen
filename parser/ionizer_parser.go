// Code generated from /Users/bneuenfeldt/Development/source/protogen/Ionizer.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Ionizer
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 62, 541,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 125, 10, 2, 12, 2, 14, 2, 128, 11, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 5, 4, 137, 10, 4, 3, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 158, 10, 7, 5, 7, 160, 10, 7, 3, 8,
	3, 8, 7, 8, 164, 10, 8, 12, 8, 14, 8, 167, 11, 8, 3, 8, 3, 8, 5, 8, 171,
	10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 177, 10, 9, 3, 10, 5, 10, 180, 10,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 190,
	10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 197, 10, 11, 12, 11,
	14, 11, 200, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 214, 10, 14, 12, 14, 14, 14, 217,
	11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 5, 15, 229, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 246,
	10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 5, 18, 269, 10, 18, 3, 19, 3, 19, 3, 19, 5, 19, 274, 10, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 7, 21, 285,
	10, 21, 12, 21, 14, 21, 288, 11, 21, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22,
	294, 10, 22, 5, 22, 296, 10, 22, 3, 23, 3, 23, 3, 23, 7, 23, 301, 10, 23,
	12, 23, 14, 23, 304, 11, 23, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 310, 10,
	24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 7, 26, 318, 10, 26, 12, 26,
	14, 26, 321, 11, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 5, 27, 328, 10,
	27, 3, 28, 3, 28, 3, 28, 5, 28, 333, 10, 28, 3, 28, 3, 28, 5, 28, 337,
	10, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 345, 10, 29, 12,
	29, 14, 29, 348, 11, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 7, 32, 362, 10, 32, 12, 32, 14, 32,
	365, 11, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	35, 3, 35, 7, 35, 377, 10, 35, 12, 35, 14, 35, 380, 11, 35, 3, 35, 3, 35,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 393,
	10, 36, 3, 37, 3, 37, 3, 37, 3, 37, 7, 37, 399, 10, 37, 12, 37, 14, 37,
	402, 11, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 5, 38, 409, 10, 38, 3,
	39, 3, 39, 3, 39, 3, 39, 5, 39, 415, 10, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 39, 5, 39, 422, 10, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 7, 39, 429,
	10, 39, 12, 39, 14, 39, 432, 11, 39, 3, 39, 3, 39, 5, 39, 436, 10, 39,
	3, 40, 3, 40, 5, 40, 440, 10, 40, 3, 40, 3, 40, 5, 40, 444, 10, 40, 3,
	40, 3, 40, 3, 40, 3, 40, 5, 40, 450, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 7, 41, 457, 10, 41, 12, 41, 14, 41, 460, 11, 41, 3, 41, 3, 41, 3,
	42, 3, 42, 3, 43, 3, 43, 5, 43, 468, 10, 43, 3, 44, 3, 44, 3, 44, 7, 44,
	473, 10, 44, 12, 44, 14, 44, 476, 11, 44, 3, 45, 3, 45, 3, 46, 5, 46, 481,
	10, 46, 3, 46, 3, 46, 3, 46, 7, 46, 486, 10, 46, 12, 46, 14, 46, 489, 11,
	46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50,
	3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 5, 53, 506, 10, 53, 3, 53, 3, 53, 3,
	53, 7, 53, 511, 10, 53, 12, 53, 14, 53, 514, 11, 53, 3, 53, 3, 53, 3, 54,
	5, 54, 519, 10, 54, 3, 54, 3, 54, 3, 54, 7, 54, 524, 10, 54, 12, 54, 14,
	54, 527, 11, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57,
	3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 2, 2, 60, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
	54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88,
	90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 2, 8,
	3, 2, 38, 39, 3, 2, 5, 6, 3, 2, 12, 23, 3, 2, 53, 54, 4, 2, 38, 39, 55,
	55, 4, 2, 3, 37, 56, 56, 2, 565, 2, 118, 3, 2, 2, 2, 4, 129, 3, 2, 2, 2,
	6, 134, 3, 2, 2, 2, 8, 141, 3, 2, 2, 2, 10, 145, 3, 2, 2, 2, 12, 159, 3,
	2, 2, 2, 14, 170, 3, 2, 2, 2, 16, 172, 3, 2, 2, 2, 18, 179, 3, 2, 2, 2,
	20, 193, 3, 2, 2, 2, 22, 201, 3, 2, 2, 2, 24, 205, 3, 2, 2, 2, 26, 207,
	3, 2, 2, 2, 28, 220, 3, 2, 2, 2, 30, 232, 3, 2, 2, 2, 32, 249, 3, 2, 2,
	2, 34, 268, 3, 2, 2, 2, 36, 270, 3, 2, 2, 2, 38, 277, 3, 2, 2, 2, 40, 281,
	3, 2, 2, 2, 42, 289, 3, 2, 2, 2, 44, 297, 3, 2, 2, 2, 46, 309, 3, 2, 2,
	2, 48, 311, 3, 2, 2, 2, 50, 315, 3, 2, 2, 2, 52, 327, 3, 2, 2, 2, 54, 329,
	3, 2, 2, 2, 56, 340, 3, 2, 2, 2, 58, 351, 3, 2, 2, 2, 60, 355, 3, 2, 2,
	2, 62, 359, 3, 2, 2, 2, 64, 368, 3, 2, 2, 2, 66, 370, 3, 2, 2, 2, 68, 374,
	3, 2, 2, 2, 70, 392, 3, 2, 2, 2, 72, 394, 3, 2, 2, 2, 74, 408, 3, 2, 2,
	2, 76, 410, 3, 2, 2, 2, 78, 449, 3, 2, 2, 2, 80, 451, 3, 2, 2, 2, 82, 463,
	3, 2, 2, 2, 84, 467, 3, 2, 2, 2, 86, 469, 3, 2, 2, 2, 88, 477, 3, 2, 2,
	2, 90, 480, 3, 2, 2, 2, 92, 492, 3, 2, 2, 2, 94, 494, 3, 2, 2, 2, 96, 496,
	3, 2, 2, 2, 98, 498, 3, 2, 2, 2, 100, 500, 3, 2, 2, 2, 102, 502, 3, 2,
	2, 2, 104, 505, 3, 2, 2, 2, 106, 518, 3, 2, 2, 2, 108, 530, 3, 2, 2, 2,
	110, 532, 3, 2, 2, 2, 112, 534, 3, 2, 2, 2, 114, 536, 3, 2, 2, 2, 116,
	538, 3, 2, 2, 2, 118, 126, 5, 4, 3, 2, 119, 125, 5, 6, 4, 2, 120, 125,
	5, 8, 5, 2, 121, 125, 5, 10, 6, 2, 122, 125, 5, 46, 24, 2, 123, 125, 5,
	82, 42, 2, 124, 119, 3, 2, 2, 2, 124, 120, 3, 2, 2, 2, 124, 121, 3, 2,
	2, 2, 124, 122, 3, 2, 2, 2, 124, 123, 3, 2, 2, 2, 125, 128, 3, 2, 2, 2,
	126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 3, 3, 2, 2, 2, 128, 126,
	3, 2, 2, 2, 129, 130, 7, 3, 2, 2, 130, 131, 7, 41, 2, 2, 131, 132, 9, 2,
	2, 2, 132, 133, 7, 40, 2, 2, 133, 5, 3, 2, 2, 2, 134, 136, 7, 4, 2, 2,
	135, 137, 9, 3, 2, 2, 136, 135, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137,
	138, 3, 2, 2, 2, 138, 139, 5, 110, 56, 2, 139, 140, 7, 40, 2, 2, 140, 7,
	3, 2, 2, 2, 141, 142, 7, 7, 2, 2, 142, 143, 5, 86, 44, 2, 143, 144, 7,
	40, 2, 2, 144, 9, 3, 2, 2, 2, 145, 146, 7, 8, 2, 2, 146, 147, 5, 12, 7,
	2, 147, 148, 7, 41, 2, 2, 148, 149, 5, 14, 8, 2, 149, 150, 7, 40, 2, 2,
	150, 11, 3, 2, 2, 2, 151, 160, 5, 86, 44, 2, 152, 153, 7, 42, 2, 2, 153,
	154, 5, 86, 44, 2, 154, 157, 7, 43, 2, 2, 155, 156, 7, 50, 2, 2, 156, 158,
	5, 86, 44, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 160, 3,
	2, 2, 2, 159, 151, 3, 2, 2, 2, 159, 152, 3, 2, 2, 2, 160, 13, 3, 2, 2,
	2, 161, 165, 7, 46, 2, 2, 162, 164, 5, 16, 9, 2, 163, 162, 3, 2, 2, 2,
	164, 167, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166,
	168, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 168, 171, 7, 47, 2, 2, 169, 171,
	5, 78, 40, 2, 170, 161, 3, 2, 2, 2, 170, 169, 3, 2, 2, 2, 171, 15, 3, 2,
	2, 2, 172, 173, 5, 94, 48, 2, 173, 174, 7, 52, 2, 2, 174, 176, 5, 14, 8,
	2, 175, 177, 7, 40, 2, 2, 176, 175, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177,
	17, 3, 2, 2, 2, 178, 180, 7, 9, 2, 2, 179, 178, 3, 2, 2, 2, 179, 180, 3,
	2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 5, 34, 18, 2, 182, 183, 5, 94,
	48, 2, 183, 184, 7, 41, 2, 2, 184, 189, 5, 24, 13, 2, 185, 186, 7, 44,
	2, 2, 186, 187, 5, 20, 11, 2, 187, 188, 7, 45, 2, 2, 188, 190, 3, 2, 2,
	2, 189, 185, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191,
	192, 7, 40, 2, 2, 192, 19, 3, 2, 2, 2, 193, 198, 5, 22, 12, 2, 194, 195,
	7, 51, 2, 2, 195, 197, 5, 22, 12, 2, 196, 194, 3, 2, 2, 2, 197, 200, 3,
	2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 21, 3, 2, 2,
	2, 200, 198, 3, 2, 2, 2, 201, 202, 5, 12, 7, 2, 202, 203, 7, 41, 2, 2,
	203, 204, 5, 78, 40, 2, 204, 23, 3, 2, 2, 2, 205, 206, 5, 108, 55, 2, 206,
	25, 3, 2, 2, 2, 207, 208, 7, 10, 2, 2, 208, 209, 5, 96, 49, 2, 209, 215,
	7, 46, 2, 2, 210, 214, 5, 10, 6, 2, 211, 214, 5, 28, 15, 2, 212, 214, 5,
	82, 42, 2, 213, 210, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 212, 3, 2,
	2, 2, 214, 217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2,
	216, 218, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 219, 7, 47, 2, 2, 219,
	27, 3, 2, 2, 2, 220, 221, 5, 34, 18, 2, 221, 222, 5, 94, 48, 2, 222, 223,
	7, 41, 2, 2, 223, 228, 5, 24, 13, 2, 224, 225, 7, 44, 2, 2, 225, 226, 5,
	20, 11, 2, 226, 227, 7, 45, 2, 2, 227, 229, 3, 2, 2, 2, 228, 224, 3, 2,
	2, 2, 228, 229, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 231, 7, 40, 2, 2,
	231, 29, 3, 2, 2, 2, 232, 233, 7, 11, 2, 2, 233, 234, 7, 48, 2, 2, 234,
	235, 5, 32, 17, 2, 235, 236, 7, 51, 2, 2, 236, 237, 5, 34, 18, 2, 237,
	238, 7, 49, 2, 2, 238, 239, 5, 98, 50, 2, 239, 240, 7, 41, 2, 2, 240, 245,
	5, 24, 13, 2, 241, 242, 7, 44, 2, 2, 242, 243, 5, 20, 11, 2, 243, 244,
	7, 45, 2, 2, 244, 246, 3, 2, 2, 2, 245, 241, 3, 2, 2, 2, 245, 246, 3, 2,
	2, 2, 246, 247, 3, 2, 2, 2, 247, 248, 7, 40, 2, 2, 248, 31, 3, 2, 2, 2,
	249, 250, 9, 4, 2, 2, 250, 33, 3, 2, 2, 2, 251, 269, 7, 24, 2, 2, 252,
	269, 7, 25, 2, 2, 253, 269, 7, 12, 2, 2, 254, 269, 7, 13, 2, 2, 255, 269,
	7, 14, 2, 2, 256, 269, 7, 15, 2, 2, 257, 269, 7, 16, 2, 2, 258, 269, 7,
	17, 2, 2, 259, 269, 7, 18, 2, 2, 260, 269, 7, 19, 2, 2, 261, 269, 7, 20,
	2, 2, 262, 269, 7, 21, 2, 2, 263, 269, 7, 22, 2, 2, 264, 269, 7, 23, 2,
	2, 265, 269, 7, 26, 2, 2, 266, 269, 5, 104, 53, 2, 267, 269, 5, 106, 54,
	2, 268, 251, 3, 2, 2, 2, 268, 252, 3, 2, 2, 2, 268, 253, 3, 2, 2, 2, 268,
	254, 3, 2, 2, 2, 268, 255, 3, 2, 2, 2, 268, 256, 3, 2, 2, 2, 268, 257,
	3, 2, 2, 2, 268, 258, 3, 2, 2, 2, 268, 259, 3, 2, 2, 2, 268, 260, 3, 2,
	2, 2, 268, 261, 3, 2, 2, 2, 268, 262, 3, 2, 2, 2, 268, 263, 3, 2, 2, 2,
	268, 264, 3, 2, 2, 2, 268, 265, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268,
	267, 3, 2, 2, 2, 269, 35, 3, 2, 2, 2, 270, 273, 7, 27, 2, 2, 271, 274,
	5, 40, 21, 2, 272, 274, 5, 44, 23, 2, 273, 271, 3, 2, 2, 2, 273, 272, 3,
	2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 276, 7, 40, 2, 2, 276, 37, 3, 2, 2,
	2, 277, 278, 7, 34, 2, 2, 278, 279, 5, 40, 21, 2, 279, 280, 7, 40, 2, 2,
	280, 39, 3, 2, 2, 2, 281, 286, 5, 42, 22, 2, 282, 283, 7, 51, 2, 2, 283,
	285, 5, 42, 22, 2, 284, 282, 3, 2, 2, 2, 285, 288, 3, 2, 2, 2, 286, 284,
	3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 41, 3, 2, 2, 2, 288, 286, 3, 2,
	2, 2, 289, 295, 5, 108, 55, 2, 290, 293, 7, 28, 2, 2, 291, 294, 5, 108,
	55, 2, 292, 294, 7, 29, 2, 2, 293, 291, 3, 2, 2, 2, 293, 292, 3, 2, 2,
	2, 294, 296, 3, 2, 2, 2, 295, 290, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296,
	43, 3, 2, 2, 2, 297, 302, 5, 110, 56, 2, 298, 299, 7, 51, 2, 2, 299, 301,
	5, 110, 56, 2, 300, 298, 3, 2, 2, 2, 301, 304, 3, 2, 2, 2, 302, 300, 3,
	2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 45, 3, 2, 2, 2, 304, 302, 3, 2, 2,
	2, 305, 310, 5, 66, 34, 2, 306, 310, 5, 60, 31, 2, 307, 310, 5, 48, 25,
	2, 308, 310, 5, 72, 37, 2, 309, 305, 3, 2, 2, 2, 309, 306, 3, 2, 2, 2,
	309, 307, 3, 2, 2, 2, 309, 308, 3, 2, 2, 2, 310, 47, 3, 2, 2, 2, 311, 312,
	7, 30, 2, 2, 312, 313, 5, 92, 47, 2, 313, 314, 5, 50, 26, 2, 314, 49, 3,
	2, 2, 2, 315, 319, 7, 46, 2, 2, 316, 318, 5, 52, 27, 2, 317, 316, 3, 2,
	2, 2, 318, 321, 3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 319, 320, 3, 2, 2, 2,
	320, 322, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2, 322, 323, 7, 47, 2, 2, 323,
	51, 3, 2, 2, 2, 324, 328, 5, 10, 6, 2, 325, 328, 5, 54, 28, 2, 326, 328,
	5, 82, 42, 2, 327, 324, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 327, 326, 3,
	2, 2, 2, 328, 53, 3, 2, 2, 2, 329, 330, 5, 84, 43, 2, 330, 332, 7, 41,
	2, 2, 331, 333, 7, 54, 2, 2, 332, 331, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2,
	333, 334, 3, 2, 2, 2, 334, 336, 5, 108, 55, 2, 335, 337, 5, 56, 29, 2,
	336, 335, 3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338,
	339, 7, 40, 2, 2, 339, 55, 3, 2, 2, 2, 340, 341, 7, 44, 2, 2, 341, 346,
	5, 58, 30, 2, 342, 343, 7, 51, 2, 2, 343, 345, 5, 58, 30, 2, 344, 342,
	3, 2, 2, 2, 345, 348, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346, 347, 3, 2,
	2, 2, 347, 349, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 349, 350, 7, 45, 2, 2,
	350, 57, 3, 2, 2, 2, 351, 352, 5, 12, 7, 2, 352, 353, 7, 41, 2, 2, 353,
	354, 5, 78, 40, 2, 354, 59, 3, 2, 2, 2, 355, 356, 7, 37, 2, 2, 356, 357,
	5, 90, 46, 2, 357, 358, 5, 62, 32, 2, 358, 61, 3, 2, 2, 2, 359, 363, 7,
	46, 2, 2, 360, 362, 5, 64, 33, 2, 361, 360, 3, 2, 2, 2, 362, 365, 3, 2,
	2, 2, 363, 361, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 366, 3, 2, 2, 2,
	365, 363, 3, 2, 2, 2, 366, 367, 7, 47, 2, 2, 367, 63, 3, 2, 2, 2, 368,
	369, 5, 18, 10, 2, 369, 65, 3, 2, 2, 2, 370, 371, 7, 31, 2, 2, 371, 372,
	5, 88, 45, 2, 372, 373, 5, 68, 35, 2, 373, 67, 3, 2, 2, 2, 374, 378, 7,
	46, 2, 2, 375, 377, 5, 70, 36, 2, 376, 375, 3, 2, 2, 2, 377, 380, 3, 2,
	2, 2, 378, 376, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 381, 3, 2, 2, 2,
	380, 378, 3, 2, 2, 2, 381, 382, 7, 47, 2, 2, 382, 69, 3, 2, 2, 2, 383,
	393, 5, 18, 10, 2, 384, 393, 5, 48, 25, 2, 385, 393, 5, 66, 34, 2, 386,
	393, 5, 10, 6, 2, 387, 393, 5, 26, 14, 2, 388, 393, 5, 30, 16, 2, 389,
	393, 5, 36, 19, 2, 390, 393, 5, 38, 20, 2, 391, 393, 5, 82, 42, 2, 392,
	383, 3, 2, 2, 2, 392, 384, 3, 2, 2, 2, 392, 385, 3, 2, 2, 2, 392, 386,
	3, 2, 2, 2, 392, 387, 3, 2, 2, 2, 392, 388, 3, 2, 2, 2, 392, 389, 3, 2,
	2, 2, 392, 390, 3, 2, 2, 2, 392, 391, 3, 2, 2, 2, 393, 71, 3, 2, 2, 2,
	394, 395, 7, 32, 2, 2, 395, 396, 5, 100, 51, 2, 396, 400, 7, 46, 2, 2,
	397, 399, 5, 74, 38, 2, 398, 397, 3, 2, 2, 2, 399, 402, 3, 2, 2, 2, 400,
	398, 3, 2, 2, 2, 400, 401, 3, 2, 2, 2, 401, 403, 3, 2, 2, 2, 402, 400,
	3, 2, 2, 2, 403, 404, 7, 47, 2, 2, 404, 73, 3, 2, 2, 2, 405, 409, 5, 10,
	6, 2, 406, 409, 5, 76, 39, 2, 407, 409, 5, 82, 42, 2, 408, 405, 3, 2, 2,
	2, 408, 406, 3, 2, 2, 2, 408, 407, 3, 2, 2, 2, 409, 75, 3, 2, 2, 2, 410,
	411, 7, 33, 2, 2, 411, 412, 5, 102, 52, 2, 412, 414, 7, 42, 2, 2, 413,
	415, 7, 35, 2, 2, 414, 413, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 416,
	3, 2, 2, 2, 416, 417, 5, 104, 53, 2, 417, 418, 7, 43, 2, 2, 418, 419, 7,
	36, 2, 2, 419, 421, 7, 42, 2, 2, 420, 422, 7, 35, 2, 2, 421, 420, 3, 2,
	2, 2, 421, 422, 3, 2, 2, 2, 422, 423, 3, 2, 2, 2, 423, 424, 5, 104, 53,
	2, 424, 435, 7, 43, 2, 2, 425, 430, 7, 46, 2, 2, 426, 429, 5, 10, 6, 2,
	427, 429, 5, 82, 42, 2, 428, 426, 3, 2, 2, 2, 428, 427, 3, 2, 2, 2, 429,
	432, 3, 2, 2, 2, 430, 428, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 431, 433,
	3, 2, 2, 2, 432, 430, 3, 2, 2, 2, 433, 436, 7, 47, 2, 2, 434, 436, 7, 40,
	2, 2, 435, 425, 3, 2, 2, 2, 435, 434, 3, 2, 2, 2, 436, 77, 3, 2, 2, 2,
	437, 450, 5, 86, 44, 2, 438, 440, 9, 5, 2, 2, 439, 438, 3, 2, 2, 2, 439,
	440, 3, 2, 2, 2, 440, 441, 3, 2, 2, 2, 441, 450, 5, 108, 55, 2, 442, 444,
	9, 5, 2, 2, 443, 442, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 445, 3, 2,
	2, 2, 445, 450, 5, 114, 58, 2, 446, 450, 5, 110, 56, 2, 447, 450, 5, 112,
	57, 2, 448, 450, 5, 80, 41, 2, 449, 437, 3, 2, 2, 2, 449, 439, 3, 2, 2,
	2, 449, 443, 3, 2, 2, 2, 449, 446, 3, 2, 2, 2, 449, 447, 3, 2, 2, 2, 449,
	448, 3, 2, 2, 2, 450, 79, 3, 2, 2, 2, 451, 458, 7, 46, 2, 2, 452, 453,
	5, 84, 43, 2, 453, 454, 7, 52, 2, 2, 454, 455, 5, 78, 40, 2, 455, 457,
	3, 2, 2, 2, 456, 452, 3, 2, 2, 2, 457, 460, 3, 2, 2, 2, 458, 456, 3, 2,
	2, 2, 458, 459, 3, 2, 2, 2, 459, 461, 3, 2, 2, 2, 460, 458, 3, 2, 2, 2,
	461, 462, 7, 47, 2, 2, 462, 81, 3, 2, 2, 2, 463, 464, 7, 40, 2, 2, 464,
	83, 3, 2, 2, 2, 465, 468, 7, 59, 2, 2, 466, 468, 5, 116, 59, 2, 467, 465,
	3, 2, 2, 2, 467, 466, 3, 2, 2, 2, 468, 85, 3, 2, 2, 2, 469, 474, 5, 84,
	43, 2, 470, 471, 7, 50, 2, 2, 471, 473, 5, 84, 43, 2, 472, 470, 3, 2, 2,
	2, 473, 476, 3, 2, 2, 2, 474, 472, 3, 2, 2, 2, 474, 475, 3, 2, 2, 2, 475,
	87, 3, 2, 2, 2, 476, 474, 3, 2, 2, 2, 477, 478, 5, 84, 43, 2, 478, 89,
	3, 2, 2, 2, 479, 481, 7, 50, 2, 2, 480, 479, 3, 2, 2, 2, 480, 481, 3, 2,
	2, 2, 481, 487, 3, 2, 2, 2, 482, 483, 5, 84, 43, 2, 483, 484, 7, 50, 2,
	2, 484, 486, 3, 2, 2, 2, 485, 482, 3, 2, 2, 2, 486, 489, 3, 2, 2, 2, 487,
	485, 3, 2, 2, 2, 487, 488, 3, 2, 2, 2, 488, 490, 3, 2, 2, 2, 489, 487,
	3, 2, 2, 2, 490, 491, 5, 84, 43, 2, 491, 91, 3, 2, 2, 2, 492, 493, 5, 84,
	43, 2, 493, 93, 3, 2, 2, 2, 494, 495, 5, 84, 43, 2, 495, 95, 3, 2, 2, 2,
	496, 497, 5, 84, 43, 2, 497, 97, 3, 2, 2, 2, 498, 499, 5, 84, 43, 2, 499,
	99, 3, 2, 2, 2, 500, 501, 5, 84, 43, 2, 501, 101, 3, 2, 2, 2, 502, 503,
	5, 84, 43, 2, 503, 103, 3, 2, 2, 2, 504, 506, 7, 50, 2, 2, 505, 504, 3,
	2, 2, 2, 505, 506, 3, 2, 2, 2, 506, 512, 3, 2, 2, 2, 507, 508, 5, 84, 43,
	2, 508, 509, 7, 50, 2, 2, 509, 511, 3, 2, 2, 2, 510, 507, 3, 2, 2, 2, 511,
	514, 3, 2, 2, 2, 512, 510, 3, 2, 2, 2, 512, 513, 3, 2, 2, 2, 513, 515,
	3, 2, 2, 2, 514, 512, 3, 2, 2, 2, 515, 516, 5, 88, 45, 2, 516, 105, 3,
	2, 2, 2, 517, 519, 7, 50, 2, 2, 518, 517, 3, 2, 2, 2, 518, 519, 3, 2, 2,
	2, 519, 525, 3, 2, 2, 2, 520, 521, 5, 84, 43, 2, 521, 522, 7, 50, 2, 2,
	522, 524, 3, 2, 2, 2, 523, 520, 3, 2, 2, 2, 524, 527, 3, 2, 2, 2, 525,
	523, 3, 2, 2, 2, 525, 526, 3, 2, 2, 2, 526, 528, 3, 2, 2, 2, 527, 525,
	3, 2, 2, 2, 528, 529, 5, 92, 47, 2, 529, 107, 3, 2, 2, 2, 530, 531, 7,
	58, 2, 2, 531, 109, 3, 2, 2, 2, 532, 533, 9, 6, 2, 2, 533, 111, 3, 2, 2,
	2, 534, 535, 7, 56, 2, 2, 535, 113, 3, 2, 2, 2, 536, 537, 7, 57, 2, 2,
	537, 115, 3, 2, 2, 2, 538, 539, 9, 7, 2, 2, 539, 117, 3, 2, 2, 2, 51, 124,
	126, 136, 157, 159, 165, 170, 176, 179, 189, 198, 213, 215, 228, 245, 268,
	273, 286, 293, 295, 302, 309, 319, 327, 332, 336, 346, 363, 378, 392, 400,
	408, 414, 421, 428, 430, 435, 439, 443, 449, 458, 467, 474, 480, 487, 505,
	512, 518, 525,
}
var literalNames = []string{
	"", "'syntax'", "'import'", "'weak'", "'public'", "'package'", "'option'",
	"'repeated'", "'oneof'", "'map'", "'int32'", "'int64'", "'uint32'", "'uint64'",
	"'sint32'", "'sint64'", "'fixed32'", "'fixed64'", "'sfixed32'", "'sfixed64'",
	"'bool'", "'string'", "'double'", "'float'", "'bytes'", "'reserved'", "'to'",
	"'max'", "'enum'", "'message'", "'service'", "'rpc'", "'extensions'", "'stream'",
	"'returns'", "'extend'", "'\"proto3\"'", "''proto3''", "';'", "'='", "'('",
	"')'", "'['", "']'", "'{'", "'}'", "'<'", "'>'", "'.'", "','", "':'", "'+'",
	"'-'",
}
var symbolicNames = []string{
	"", "SYNTAX", "IMPORT", "WEAK", "PUBLIC", "PACKAGE", "OPTION", "REPEATED",
	"ONEOF", "MAP", "INT32", "INT64", "UINT32", "UINT64", "SINT32", "SINT64",
	"FIXED32", "FIXED64", "SFIXED32", "SFIXED64", "BOOL", "STRING", "DOUBLE",
	"FLOAT", "BYTES", "RESERVED", "TO", "MAX", "ENUM", "MESSAGE", "SERVICE",
	"RPC", "EXTENSIONS", "STREAM", "RETURNS", "EXTEND", "PROTO3_LIT_SINGLE",
	"PROTO3_LIT_DOBULE", "SEMI", "EQ", "LP", "RP", "LB", "RB", "LC", "RC",
	"LT", "GT", "DOT", "COMMA", "COLON", "PLUS", "MINUS", "STR_LIT", "BOOL_LIT",
	"FLOAT_LIT", "INT_LIT", "IDENTIFIER", "WS", "LINE_COMMENT", "COMMENT",
}

var ruleNames = []string{
	"proto", "syntax", "importStatement", "packageStatement", "optionStatement",
	"optionName", "optionBody", "optionElement", "field", "fieldOptions", "fieldOption",
	"fieldNumber", "oneof", "oneofField", "mapField", "keyType", "type_", "reserved",
	"extensionStatement", "ranges", "range_", "reservedFieldNames", "topLevelDef",
	"enumDef", "enumBody", "enumElement", "enumField", "enumValueOptions",
	"enumValueOption", "extendDef", "extendBody", "extendElement", "messageDef",
	"messageBody", "messageElement", "serviceDef", "serviceElement", "rpc",
	"constant", "blockLit", "emptyStatement_", "ident", "fullIdent", "messageName",
	"extendName", "enumName", "fieldName", "oneofName", "mapName", "serviceName",
	"rpcName", "messageType", "enumType", "intLit", "strLit", "boolLit", "floatLit",
	"keywords",
}

type IonizerParser struct {
	*antlr.BaseParser
}

// NewIonizerParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *IonizerParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewIonizerParser(input antlr.TokenStream) *IonizerParser {
	this := new(IonizerParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Ionizer.g4"

	return this
}

// IonizerParser tokens.
const (
	IonizerParserEOF               = antlr.TokenEOF
	IonizerParserSYNTAX            = 1
	IonizerParserIMPORT            = 2
	IonizerParserWEAK              = 3
	IonizerParserPUBLIC            = 4
	IonizerParserPACKAGE           = 5
	IonizerParserOPTION            = 6
	IonizerParserREPEATED          = 7
	IonizerParserONEOF             = 8
	IonizerParserMAP               = 9
	IonizerParserINT32             = 10
	IonizerParserINT64             = 11
	IonizerParserUINT32            = 12
	IonizerParserUINT64            = 13
	IonizerParserSINT32            = 14
	IonizerParserSINT64            = 15
	IonizerParserFIXED32           = 16
	IonizerParserFIXED64           = 17
	IonizerParserSFIXED32          = 18
	IonizerParserSFIXED64          = 19
	IonizerParserBOOL              = 20
	IonizerParserSTRING            = 21
	IonizerParserDOUBLE            = 22
	IonizerParserFLOAT             = 23
	IonizerParserBYTES             = 24
	IonizerParserRESERVED          = 25
	IonizerParserTO                = 26
	IonizerParserMAX               = 27
	IonizerParserENUM              = 28
	IonizerParserMESSAGE           = 29
	IonizerParserSERVICE           = 30
	IonizerParserRPC               = 31
	IonizerParserEXTENSIONS        = 32
	IonizerParserSTREAM            = 33
	IonizerParserRETURNS           = 34
	IonizerParserEXTEND            = 35
	IonizerParserPROTO3_LIT_SINGLE = 36
	IonizerParserPROTO3_LIT_DOBULE = 37
	IonizerParserSEMI              = 38
	IonizerParserEQ                = 39
	IonizerParserLP                = 40
	IonizerParserRP                = 41
	IonizerParserLB                = 42
	IonizerParserRB                = 43
	IonizerParserLC                = 44
	IonizerParserRC                = 45
	IonizerParserLT                = 46
	IonizerParserGT                = 47
	IonizerParserDOT               = 48
	IonizerParserCOMMA             = 49
	IonizerParserCOLON             = 50
	IonizerParserPLUS              = 51
	IonizerParserMINUS             = 52
	IonizerParserSTR_LIT           = 53
	IonizerParserBOOL_LIT          = 54
	IonizerParserFLOAT_LIT         = 55
	IonizerParserINT_LIT           = 56
	IonizerParserIDENTIFIER        = 57
	IonizerParserWS                = 58
	IonizerParserLINE_COMMENT      = 59
	IonizerParserCOMMENT           = 60
)

// IonizerParser rules.
const (
	IonizerParserRULE_proto              = 0
	IonizerParserRULE_syntax             = 1
	IonizerParserRULE_importStatement    = 2
	IonizerParserRULE_packageStatement   = 3
	IonizerParserRULE_optionStatement    = 4
	IonizerParserRULE_optionName         = 5
	IonizerParserRULE_optionBody         = 6
	IonizerParserRULE_optionElement      = 7
	IonizerParserRULE_field              = 8
	IonizerParserRULE_fieldOptions       = 9
	IonizerParserRULE_fieldOption        = 10
	IonizerParserRULE_fieldNumber        = 11
	IonizerParserRULE_oneof              = 12
	IonizerParserRULE_oneofField         = 13
	IonizerParserRULE_mapField           = 14
	IonizerParserRULE_keyType            = 15
	IonizerParserRULE_type_              = 16
	IonizerParserRULE_reserved           = 17
	IonizerParserRULE_extensionStatement = 18
	IonizerParserRULE_ranges             = 19
	IonizerParserRULE_range_             = 20
	IonizerParserRULE_reservedFieldNames = 21
	IonizerParserRULE_topLevelDef        = 22
	IonizerParserRULE_enumDef            = 23
	IonizerParserRULE_enumBody           = 24
	IonizerParserRULE_enumElement        = 25
	IonizerParserRULE_enumField          = 26
	IonizerParserRULE_enumValueOptions   = 27
	IonizerParserRULE_enumValueOption    = 28
	IonizerParserRULE_extendDef          = 29
	IonizerParserRULE_extendBody         = 30
	IonizerParserRULE_extendElement      = 31
	IonizerParserRULE_messageDef         = 32
	IonizerParserRULE_messageBody        = 33
	IonizerParserRULE_messageElement     = 34
	IonizerParserRULE_serviceDef         = 35
	IonizerParserRULE_serviceElement     = 36
	IonizerParserRULE_rpc                = 37
	IonizerParserRULE_constant           = 38
	IonizerParserRULE_blockLit           = 39
	IonizerParserRULE_emptyStatement_    = 40
	IonizerParserRULE_ident              = 41
	IonizerParserRULE_fullIdent          = 42
	IonizerParserRULE_messageName        = 43
	IonizerParserRULE_extendName         = 44
	IonizerParserRULE_enumName           = 45
	IonizerParserRULE_fieldName          = 46
	IonizerParserRULE_oneofName          = 47
	IonizerParserRULE_mapName            = 48
	IonizerParserRULE_serviceName        = 49
	IonizerParserRULE_rpcName            = 50
	IonizerParserRULE_messageType        = 51
	IonizerParserRULE_enumType           = 52
	IonizerParserRULE_intLit             = 53
	IonizerParserRULE_strLit             = 54
	IonizerParserRULE_boolLit            = 55
	IonizerParserRULE_floatLit           = 56
	IonizerParserRULE_keywords           = 57
)

// IProtoContext is an interface to support dynamic dispatch.
type IProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProtoContext differentiates from other interfaces.
	IsProtoContext()
}

type ProtoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoContext() *ProtoContext {
	var p = new(ProtoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_proto
	return p
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) Syntax() ISyntaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ProtoContext) AllImportStatement() []IImportStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportStatementContext)(nil)).Elem())
	var tst = make([]IImportStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportStatementContext)
		}
	}

	return tst
}

func (s *ProtoContext) ImportStatement(i int) IImportStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *ProtoContext) AllPackageStatement() []IPackageStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPackageStatementContext)(nil)).Elem())
	var tst = make([]IPackageStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPackageStatementContext)
		}
	}

	return tst
}

func (s *ProtoContext) PackageStatement(i int) IPackageStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPackageStatementContext)
}

func (s *ProtoContext) AllOptionStatement() []IOptionStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem())
	var tst = make([]IOptionStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionStatementContext)
		}
	}

	return tst
}

func (s *ProtoContext) OptionStatement(i int) IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *ProtoContext) AllTopLevelDef() []ITopLevelDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopLevelDefContext)(nil)).Elem())
	var tst = make([]ITopLevelDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopLevelDefContext)
		}
	}

	return tst
}

func (s *ProtoContext) TopLevelDef(i int) ITopLevelDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopLevelDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopLevelDefContext)
}

func (s *ProtoContext) AllEmptyStatement_() []IEmptyStatement_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem())
	var tst = make([]IEmptyStatement_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatement_Context)
		}
	}

	return tst
}

func (s *ProtoContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *ProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterProto(s)
	}
}

func (s *ProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitProto(s)
	}
}

func (s *ProtoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitProto(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IonizerParserRULE_proto)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Syntax()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IonizerParserIMPORT)|(1<<IonizerParserPACKAGE)|(1<<IonizerParserOPTION)|(1<<IonizerParserENUM)|(1<<IonizerParserMESSAGE)|(1<<IonizerParserSERVICE))) != 0) || _la == IonizerParserEXTEND || _la == IonizerParserSEMI {
		p.SetState(122)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IonizerParserIMPORT:
			{
				p.SetState(117)
				p.ImportStatement()
			}

		case IonizerParserPACKAGE:
			{
				p.SetState(118)
				p.PackageStatement()
			}

		case IonizerParserOPTION:
			{
				p.SetState(119)
				p.OptionStatement()
			}

		case IonizerParserENUM, IonizerParserMESSAGE, IonizerParserSERVICE, IonizerParserEXTEND:
			{
				p.SetState(120)
				p.TopLevelDef()
			}

		case IonizerParserSEMI:
			{
				p.SetState(121)
				p.EmptyStatement_()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISyntaxContext is an interface to support dynamic dispatch.
type ISyntaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSyntaxContext differentiates from other interfaces.
	IsSyntaxContext()
}

type SyntaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntaxContext() *SyntaxContext {
	var p = new(SyntaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_syntax
	return p
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(IonizerParserSYNTAX, 0)
}

func (s *SyntaxContext) EQ() antlr.TerminalNode {
	return s.GetToken(IonizerParserEQ, 0)
}

func (s *SyntaxContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *SyntaxContext) PROTO3_LIT_SINGLE() antlr.TerminalNode {
	return s.GetToken(IonizerParserPROTO3_LIT_SINGLE, 0)
}

func (s *SyntaxContext) PROTO3_LIT_DOBULE() antlr.TerminalNode {
	return s.GetToken(IonizerParserPROTO3_LIT_DOBULE, 0)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterSyntax(s)
	}
}

func (s *SyntaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitSyntax(s)
	}
}

func (s *SyntaxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitSyntax(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Syntax() (localctx ISyntaxContext) {
	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IonizerParserRULE_syntax)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(IonizerParserSYNTAX)
	}
	{
		p.SetState(128)
		p.Match(IonizerParserEQ)
	}
	{
		p.SetState(129)
		_la = p.GetTokenStream().LA(1)

		if !(_la == IonizerParserPROTO3_LIT_SINGLE || _la == IonizerParserPROTO3_LIT_DOBULE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(130)
		p.Match(IonizerParserSEMI)
	}

	return localctx
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_importStatement
	return p
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(IonizerParserIMPORT, 0)
}

func (s *ImportStatementContext) StrLit() IStrLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *ImportStatementContext) WEAK() antlr.TerminalNode {
	return s.GetToken(IonizerParserWEAK, 0)
}

func (s *ImportStatementContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(IonizerParserPUBLIC, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (s *ImportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitImportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, IonizerParserRULE_importStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(IonizerParserIMPORT)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IonizerParserWEAK || _la == IonizerParserPUBLIC {
		{
			p.SetState(133)
			_la = p.GetTokenStream().LA(1)

			if !(_la == IonizerParserWEAK || _la == IonizerParserPUBLIC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(136)
		p.StrLit()
	}
	{
		p.SetState(137)
		p.Match(IonizerParserSEMI)
	}

	return localctx
}

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_packageStatement
	return p
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(IonizerParserPACKAGE, 0)
}

func (s *PackageStatementContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (s *PackageStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitPackageStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) PackageStatement() (localctx IPackageStatementContext) {
	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, IonizerParserRULE_packageStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(IonizerParserPACKAGE)
	}
	{
		p.SetState(140)
		p.FullIdent()
	}
	{
		p.SetState(141)
		p.Match(IonizerParserSEMI)
	}

	return localctx
}

// IOptionStatementContext is an interface to support dynamic dispatch.
type IOptionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionStatementContext differentiates from other interfaces.
	IsOptionStatementContext()
}

type OptionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionStatementContext() *OptionStatementContext {
	var p = new(OptionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_optionStatement
	return p
}

func (*OptionStatementContext) IsOptionStatementContext() {}

func NewOptionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionStatementContext {
	var p = new(OptionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_optionStatement

	return p
}

func (s *OptionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionStatementContext) OPTION() antlr.TerminalNode {
	return s.GetToken(IonizerParserOPTION, 0)
}

func (s *OptionStatementContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(IonizerParserEQ, 0)
}

func (s *OptionStatementContext) OptionBody() IOptionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionBodyContext)
}

func (s *OptionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *OptionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterOptionStatement(s)
	}
}

func (s *OptionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitOptionStatement(s)
	}
}

func (s *OptionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitOptionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) OptionStatement() (localctx IOptionStatementContext) {
	localctx = NewOptionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, IonizerParserRULE_optionStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(IonizerParserOPTION)
	}
	{
		p.SetState(144)
		p.OptionName()
	}
	{
		p.SetState(145)
		p.Match(IonizerParserEQ)
	}
	{
		p.SetState(146)
		p.OptionBody()
	}
	{
		p.SetState(147)
		p.Match(IonizerParserSEMI)
	}

	return localctx
}

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_optionName
	return p
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) AllFullIdent() []IFullIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFullIdentContext)(nil)).Elem())
	var tst = make([]IFullIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFullIdentContext)
		}
	}

	return tst
}

func (s *OptionNameContext) FullIdent(i int) IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *OptionNameContext) LP() antlr.TerminalNode {
	return s.GetToken(IonizerParserLP, 0)
}

func (s *OptionNameContext) RP() antlr.TerminalNode {
	return s.GetToken(IonizerParserRP, 0)
}

func (s *OptionNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(IonizerParserDOT, 0)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterOptionName(s)
	}
}

func (s *OptionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitOptionName(s)
	}
}

func (s *OptionNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitOptionName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, IonizerParserRULE_optionName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IonizerParserSYNTAX, IonizerParserIMPORT, IonizerParserWEAK, IonizerParserPUBLIC, IonizerParserPACKAGE, IonizerParserOPTION, IonizerParserREPEATED, IonizerParserONEOF, IonizerParserMAP, IonizerParserINT32, IonizerParserINT64, IonizerParserUINT32, IonizerParserUINT64, IonizerParserSINT32, IonizerParserSINT64, IonizerParserFIXED32, IonizerParserFIXED64, IonizerParserSFIXED32, IonizerParserSFIXED64, IonizerParserBOOL, IonizerParserSTRING, IonizerParserDOUBLE, IonizerParserFLOAT, IonizerParserBYTES, IonizerParserRESERVED, IonizerParserTO, IonizerParserMAX, IonizerParserENUM, IonizerParserMESSAGE, IonizerParserSERVICE, IonizerParserRPC, IonizerParserEXTENSIONS, IonizerParserSTREAM, IonizerParserRETURNS, IonizerParserEXTEND, IonizerParserBOOL_LIT, IonizerParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.FullIdent()
		}

	case IonizerParserLP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Match(IonizerParserLP)
		}
		{
			p.SetState(151)
			p.FullIdent()
		}
		{
			p.SetState(152)
			p.Match(IonizerParserRP)
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IonizerParserDOT {
			{
				p.SetState(153)
				p.Match(IonizerParserDOT)
			}
			{
				p.SetState(154)
				p.FullIdent()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOptionBodyContext is an interface to support dynamic dispatch.
type IOptionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionBodyContext differentiates from other interfaces.
	IsOptionBodyContext()
}

type OptionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionBodyContext() *OptionBodyContext {
	var p = new(OptionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_optionBody
	return p
}

func (*OptionBodyContext) IsOptionBodyContext() {}

func NewOptionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionBodyContext {
	var p = new(OptionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_optionBody

	return p
}

func (s *OptionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(IonizerParserLC, 0)
}

func (s *OptionBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(IonizerParserRC, 0)
}

func (s *OptionBodyContext) AllOptionElement() []IOptionElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionElementContext)(nil)).Elem())
	var tst = make([]IOptionElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionElementContext)
		}
	}

	return tst
}

func (s *OptionBodyContext) OptionElement(i int) IOptionElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionElementContext)
}

func (s *OptionBodyContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterOptionBody(s)
	}
}

func (s *OptionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitOptionBody(s)
	}
}

func (s *OptionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitOptionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) OptionBody() (localctx IOptionBodyContext) {
	localctx = NewOptionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, IonizerParserRULE_optionBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.Match(IonizerParserLC)
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IonizerParserSYNTAX)|(1<<IonizerParserIMPORT)|(1<<IonizerParserWEAK)|(1<<IonizerParserPUBLIC)|(1<<IonizerParserPACKAGE)|(1<<IonizerParserOPTION)|(1<<IonizerParserREPEATED)|(1<<IonizerParserONEOF)|(1<<IonizerParserMAP)|(1<<IonizerParserINT32)|(1<<IonizerParserINT64)|(1<<IonizerParserUINT32)|(1<<IonizerParserUINT64)|(1<<IonizerParserSINT32)|(1<<IonizerParserSINT64)|(1<<IonizerParserFIXED32)|(1<<IonizerParserFIXED64)|(1<<IonizerParserSFIXED32)|(1<<IonizerParserSFIXED64)|(1<<IonizerParserBOOL)|(1<<IonizerParserSTRING)|(1<<IonizerParserDOUBLE)|(1<<IonizerParserFLOAT)|(1<<IonizerParserBYTES)|(1<<IonizerParserRESERVED)|(1<<IonizerParserTO)|(1<<IonizerParserMAX)|(1<<IonizerParserENUM)|(1<<IonizerParserMESSAGE)|(1<<IonizerParserSERVICE)|(1<<IonizerParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IonizerParserEXTENSIONS-32))|(1<<(IonizerParserSTREAM-32))|(1<<(IonizerParserRETURNS-32))|(1<<(IonizerParserEXTEND-32))|(1<<(IonizerParserBOOL_LIT-32))|(1<<(IonizerParserIDENTIFIER-32)))) != 0) {
			{
				p.SetState(160)
				p.OptionElement()
			}

			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(166)
			p.Match(IonizerParserRC)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Constant()
		}

	}

	return localctx
}

// IOptionElementContext is an interface to support dynamic dispatch.
type IOptionElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionElementContext differentiates from other interfaces.
	IsOptionElementContext()
}

type OptionElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionElementContext() *OptionElementContext {
	var p = new(OptionElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_optionElement
	return p
}

func (*OptionElementContext) IsOptionElementContext() {}

func NewOptionElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionElementContext {
	var p = new(OptionElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_optionElement

	return p
}

func (s *OptionElementContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionElementContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *OptionElementContext) COLON() antlr.TerminalNode {
	return s.GetToken(IonizerParserCOLON, 0)
}

func (s *OptionElementContext) OptionBody() IOptionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionBodyContext)
}

func (s *OptionElementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *OptionElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterOptionElement(s)
	}
}

func (s *OptionElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitOptionElement(s)
	}
}

func (s *OptionElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitOptionElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) OptionElement() (localctx IOptionElementContext) {
	localctx = NewOptionElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, IonizerParserRULE_optionElement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.FieldName()
	}
	{
		p.SetState(171)
		p.Match(IonizerParserCOLON)
	}
	{
		p.SetState(172)
		p.OptionBody()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IonizerParserSEMI {
		{
			p.SetState(173)
			p.Match(IonizerParserSEMI)
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(IonizerParserEQ, 0)
}

func (s *FieldContext) FieldNumber() IFieldNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *FieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *FieldContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(IonizerParserREPEATED, 0)
}

func (s *FieldContext) LB() antlr.TerminalNode {
	return s.GetToken(IonizerParserLB, 0)
}

func (s *FieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldContext) RB() antlr.TerminalNode {
	return s.GetToken(IonizerParserRB, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, IonizerParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(177)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(176)
			p.Match(IonizerParserREPEATED)
		}

	}
	{
		p.SetState(179)
		p.Type_()
	}
	{
		p.SetState(180)
		p.FieldName()
	}
	{
		p.SetState(181)
		p.Match(IonizerParserEQ)
	}
	{
		p.SetState(182)
		p.FieldNumber()
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IonizerParserLB {
		{
			p.SetState(183)
			p.Match(IonizerParserLB)
		}
		{
			p.SetState(184)
			p.FieldOptions()
		}
		{
			p.SetState(185)
			p.Match(IonizerParserRB)
		}

	}
	{
		p.SetState(189)
		p.Match(IonizerParserSEMI)
	}

	return localctx
}

// IFieldOptionsContext is an interface to support dynamic dispatch.
type IFieldOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptionsContext differentiates from other interfaces.
	IsFieldOptionsContext()
}

type FieldOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionsContext() *FieldOptionsContext {
	var p = new(FieldOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_fieldOptions
	return p
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_fieldOptions

	return p
}

func (s *FieldOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionsContext) AllFieldOption() []IFieldOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldOptionContext)(nil)).Elem())
	var tst = make([]IFieldOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldOptionContext)
		}
	}

	return tst
}

func (s *FieldOptionsContext) FieldOption(i int) IFieldOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionContext)
}

func (s *FieldOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserCOMMA)
}

func (s *FieldOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserCOMMA, i)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterFieldOptions(s)
	}
}

func (s *FieldOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitFieldOptions(s)
	}
}

func (s *FieldOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitFieldOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) FieldOptions() (localctx IFieldOptionsContext) {
	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, IonizerParserRULE_fieldOptions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.FieldOption()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IonizerParserCOMMA {
		{
			p.SetState(192)
			p.Match(IonizerParserCOMMA)
		}
		{
			p.SetState(193)
			p.FieldOption()
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldOptionContext is an interface to support dynamic dispatch.
type IFieldOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptionContext differentiates from other interfaces.
	IsFieldOptionContext()
}

type FieldOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionContext() *FieldOptionContext {
	var p = new(FieldOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_fieldOption
	return p
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_fieldOption

	return p
}

func (s *FieldOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *FieldOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(IonizerParserEQ, 0)
}

func (s *FieldOptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FieldOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterFieldOption(s)
	}
}

func (s *FieldOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitFieldOption(s)
	}
}

func (s *FieldOptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitFieldOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) FieldOption() (localctx IFieldOptionContext) {
	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, IonizerParserRULE_fieldOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.OptionName()
	}
	{
		p.SetState(200)
		p.Match(IonizerParserEQ)
	}
	{
		p.SetState(201)
		p.Constant()
	}

	return localctx
}

// IFieldNumberContext is an interface to support dynamic dispatch.
type IFieldNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNumberContext differentiates from other interfaces.
	IsFieldNumberContext()
}

type FieldNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNumberContext() *FieldNumberContext {
	var p = new(FieldNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_fieldNumber
	return p
}

func (*FieldNumberContext) IsFieldNumberContext() {}

func NewFieldNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNumberContext {
	var p = new(FieldNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_fieldNumber

	return p
}

func (s *FieldNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNumberContext) IntLit() IIntLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *FieldNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterFieldNumber(s)
	}
}

func (s *FieldNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitFieldNumber(s)
	}
}

func (s *FieldNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitFieldNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) FieldNumber() (localctx IFieldNumberContext) {
	localctx = NewFieldNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, IonizerParserRULE_fieldNumber)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.IntLit()
	}

	return localctx
}

// IOneofContext is an interface to support dynamic dispatch.
type IOneofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofContext differentiates from other interfaces.
	IsOneofContext()
}

type OneofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofContext() *OneofContext {
	var p = new(OneofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_oneof
	return p
}

func (*OneofContext) IsOneofContext() {}

func NewOneofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofContext {
	var p = new(OneofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_oneof

	return p
}

func (s *OneofContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(IonizerParserONEOF, 0)
}

func (s *OneofContext) OneofName() IOneofNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOneofNameContext)
}

func (s *OneofContext) LC() antlr.TerminalNode {
	return s.GetToken(IonizerParserLC, 0)
}

func (s *OneofContext) RC() antlr.TerminalNode {
	return s.GetToken(IonizerParserRC, 0)
}

func (s *OneofContext) AllOptionStatement() []IOptionStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem())
	var tst = make([]IOptionStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionStatementContext)
		}
	}

	return tst
}

func (s *OneofContext) OptionStatement(i int) IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *OneofContext) AllOneofField() []IOneofFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOneofFieldContext)(nil)).Elem())
	var tst = make([]IOneofFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOneofFieldContext)
		}
	}

	return tst
}

func (s *OneofContext) OneofField(i int) IOneofFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOneofFieldContext)
}

func (s *OneofContext) AllEmptyStatement_() []IEmptyStatement_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem())
	var tst = make([]IEmptyStatement_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatement_Context)
		}
	}

	return tst
}

func (s *OneofContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *OneofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterOneof(s)
	}
}

func (s *OneofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitOneof(s)
	}
}

func (s *OneofContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitOneof(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Oneof() (localctx IOneofContext) {
	localctx = NewOneofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, IonizerParserRULE_oneof)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(IonizerParserONEOF)
	}
	{
		p.SetState(206)
		p.OneofName()
	}
	{
		p.SetState(207)
		p.Match(IonizerParserLC)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IonizerParserSYNTAX)|(1<<IonizerParserIMPORT)|(1<<IonizerParserWEAK)|(1<<IonizerParserPUBLIC)|(1<<IonizerParserPACKAGE)|(1<<IonizerParserOPTION)|(1<<IonizerParserREPEATED)|(1<<IonizerParserONEOF)|(1<<IonizerParserMAP)|(1<<IonizerParserINT32)|(1<<IonizerParserINT64)|(1<<IonizerParserUINT32)|(1<<IonizerParserUINT64)|(1<<IonizerParserSINT32)|(1<<IonizerParserSINT64)|(1<<IonizerParserFIXED32)|(1<<IonizerParserFIXED64)|(1<<IonizerParserSFIXED32)|(1<<IonizerParserSFIXED64)|(1<<IonizerParserBOOL)|(1<<IonizerParserSTRING)|(1<<IonizerParserDOUBLE)|(1<<IonizerParserFLOAT)|(1<<IonizerParserBYTES)|(1<<IonizerParserRESERVED)|(1<<IonizerParserTO)|(1<<IonizerParserMAX)|(1<<IonizerParserENUM)|(1<<IonizerParserMESSAGE)|(1<<IonizerParserSERVICE)|(1<<IonizerParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IonizerParserEXTENSIONS-32))|(1<<(IonizerParserSTREAM-32))|(1<<(IonizerParserRETURNS-32))|(1<<(IonizerParserEXTEND-32))|(1<<(IonizerParserSEMI-32))|(1<<(IonizerParserDOT-32))|(1<<(IonizerParserBOOL_LIT-32))|(1<<(IonizerParserIDENTIFIER-32)))) != 0) {
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(208)
				p.OptionStatement()
			}

		case 2:
			{
				p.SetState(209)
				p.OneofField()
			}

		case 3:
			{
				p.SetState(210)
				p.EmptyStatement_()
			}

		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(IonizerParserRC)
	}

	return localctx
}

// IOneofFieldContext is an interface to support dynamic dispatch.
type IOneofFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofFieldContext differentiates from other interfaces.
	IsOneofFieldContext()
}

type OneofFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofFieldContext() *OneofFieldContext {
	var p = new(OneofFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_oneofField
	return p
}

func (*OneofFieldContext) IsOneofFieldContext() {}

func NewOneofFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofFieldContext {
	var p = new(OneofFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_oneofField

	return p
}

func (s *OneofFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofFieldContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *OneofFieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *OneofFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(IonizerParserEQ, 0)
}

func (s *OneofFieldContext) FieldNumber() IFieldNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *OneofFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *OneofFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(IonizerParserLB, 0)
}

func (s *OneofFieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *OneofFieldContext) RB() antlr.TerminalNode {
	return s.GetToken(IonizerParserRB, 0)
}

func (s *OneofFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterOneofField(s)
	}
}

func (s *OneofFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitOneofField(s)
	}
}

func (s *OneofFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitOneofField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) OneofField() (localctx IOneofFieldContext) {
	localctx = NewOneofFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, IonizerParserRULE_oneofField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Type_()
	}
	{
		p.SetState(219)
		p.FieldName()
	}
	{
		p.SetState(220)
		p.Match(IonizerParserEQ)
	}
	{
		p.SetState(221)
		p.FieldNumber()
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IonizerParserLB {
		{
			p.SetState(222)
			p.Match(IonizerParserLB)
		}
		{
			p.SetState(223)
			p.FieldOptions()
		}
		{
			p.SetState(224)
			p.Match(IonizerParserRB)
		}

	}
	{
		p.SetState(228)
		p.Match(IonizerParserSEMI)
	}

	return localctx
}

// IMapFieldContext is an interface to support dynamic dispatch.
type IMapFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapFieldContext differentiates from other interfaces.
	IsMapFieldContext()
}

type MapFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapFieldContext() *MapFieldContext {
	var p = new(MapFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_mapField
	return p
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) MAP() antlr.TerminalNode {
	return s.GetToken(IonizerParserMAP, 0)
}

func (s *MapFieldContext) LT() antlr.TerminalNode {
	return s.GetToken(IonizerParserLT, 0)
}

func (s *MapFieldContext) KeyType() IKeyTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *MapFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(IonizerParserCOMMA, 0)
}

func (s *MapFieldContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *MapFieldContext) GT() antlr.TerminalNode {
	return s.GetToken(IonizerParserGT, 0)
}

func (s *MapFieldContext) MapName() IMapNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapNameContext)
}

func (s *MapFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(IonizerParserEQ, 0)
}

func (s *MapFieldContext) FieldNumber() IFieldNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *MapFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *MapFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(IonizerParserLB, 0)
}

func (s *MapFieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *MapFieldContext) RB() antlr.TerminalNode {
	return s.GetToken(IonizerParserRB, 0)
}

func (s *MapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterMapField(s)
	}
}

func (s *MapFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitMapField(s)
	}
}

func (s *MapFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitMapField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) MapField() (localctx IMapFieldContext) {
	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, IonizerParserRULE_mapField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(IonizerParserMAP)
	}
	{
		p.SetState(231)
		p.Match(IonizerParserLT)
	}
	{
		p.SetState(232)
		p.KeyType()
	}
	{
		p.SetState(233)
		p.Match(IonizerParserCOMMA)
	}
	{
		p.SetState(234)
		p.Type_()
	}
	{
		p.SetState(235)
		p.Match(IonizerParserGT)
	}
	{
		p.SetState(236)
		p.MapName()
	}
	{
		p.SetState(237)
		p.Match(IonizerParserEQ)
	}
	{
		p.SetState(238)
		p.FieldNumber()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IonizerParserLB {
		{
			p.SetState(239)
			p.Match(IonizerParserLB)
		}
		{
			p.SetState(240)
			p.FieldOptions()
		}
		{
			p.SetState(241)
			p.Match(IonizerParserRB)
		}

	}
	{
		p.SetState(245)
		p.Match(IonizerParserSEMI)
	}

	return localctx
}

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyTypeContext differentiates from other interfaces.
	IsKeyTypeContext()
}

type KeyTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyTypeContext() *KeyTypeContext {
	var p = new(KeyTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_keyType
	return p
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(IonizerParserINT32, 0)
}

func (s *KeyTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(IonizerParserINT64, 0)
}

func (s *KeyTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(IonizerParserUINT32, 0)
}

func (s *KeyTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(IonizerParserUINT64, 0)
}

func (s *KeyTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(IonizerParserSINT32, 0)
}

func (s *KeyTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(IonizerParserSINT64, 0)
}

func (s *KeyTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(IonizerParserFIXED32, 0)
}

func (s *KeyTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(IonizerParserFIXED64, 0)
}

func (s *KeyTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(IonizerParserSFIXED32, 0)
}

func (s *KeyTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(IonizerParserSFIXED64, 0)
}

func (s *KeyTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(IonizerParserBOOL, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(IonizerParserSTRING, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (s *KeyTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitKeyType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) KeyType() (localctx IKeyTypeContext) {
	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, IonizerParserRULE_keyType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IonizerParserINT32)|(1<<IonizerParserINT64)|(1<<IonizerParserUINT32)|(1<<IonizerParserUINT64)|(1<<IonizerParserSINT32)|(1<<IonizerParserSINT64)|(1<<IonizerParserFIXED32)|(1<<IonizerParserFIXED64)|(1<<IonizerParserSFIXED32)|(1<<IonizerParserSFIXED64)|(1<<IonizerParserBOOL)|(1<<IonizerParserSTRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) DOUBLE() antlr.TerminalNode {
	return s.GetToken(IonizerParserDOUBLE, 0)
}

func (s *Type_Context) FLOAT() antlr.TerminalNode {
	return s.GetToken(IonizerParserFLOAT, 0)
}

func (s *Type_Context) INT32() antlr.TerminalNode {
	return s.GetToken(IonizerParserINT32, 0)
}

func (s *Type_Context) INT64() antlr.TerminalNode {
	return s.GetToken(IonizerParserINT64, 0)
}

func (s *Type_Context) UINT32() antlr.TerminalNode {
	return s.GetToken(IonizerParserUINT32, 0)
}

func (s *Type_Context) UINT64() antlr.TerminalNode {
	return s.GetToken(IonizerParserUINT64, 0)
}

func (s *Type_Context) SINT32() antlr.TerminalNode {
	return s.GetToken(IonizerParserSINT32, 0)
}

func (s *Type_Context) SINT64() antlr.TerminalNode {
	return s.GetToken(IonizerParserSINT64, 0)
}

func (s *Type_Context) FIXED32() antlr.TerminalNode {
	return s.GetToken(IonizerParserFIXED32, 0)
}

func (s *Type_Context) FIXED64() antlr.TerminalNode {
	return s.GetToken(IonizerParserFIXED64, 0)
}

func (s *Type_Context) SFIXED32() antlr.TerminalNode {
	return s.GetToken(IonizerParserSFIXED32, 0)
}

func (s *Type_Context) SFIXED64() antlr.TerminalNode {
	return s.GetToken(IonizerParserSFIXED64, 0)
}

func (s *Type_Context) BOOL() antlr.TerminalNode {
	return s.GetToken(IonizerParserBOOL, 0)
}

func (s *Type_Context) STRING() antlr.TerminalNode {
	return s.GetToken(IonizerParserSTRING, 0)
}

func (s *Type_Context) BYTES() antlr.TerminalNode {
	return s.GetToken(IonizerParserBYTES, 0)
}

func (s *Type_Context) MessageType() IMessageTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *Type_Context) EnumType() IEnumTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitType_(s)
	}
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, IonizerParserRULE_type_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.Match(IonizerParserDOUBLE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Match(IonizerParserFLOAT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(251)
			p.Match(IonizerParserINT32)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(252)
			p.Match(IonizerParserINT64)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(253)
			p.Match(IonizerParserUINT32)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(254)
			p.Match(IonizerParserUINT64)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(255)
			p.Match(IonizerParserSINT32)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(256)
			p.Match(IonizerParserSINT64)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(257)
			p.Match(IonizerParserFIXED32)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(258)
			p.Match(IonizerParserFIXED64)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(259)
			p.Match(IonizerParserSFIXED32)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(260)
			p.Match(IonizerParserSFIXED64)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(261)
			p.Match(IonizerParserBOOL)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(262)
			p.Match(IonizerParserSTRING)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(263)
			p.Match(IonizerParserBYTES)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(264)
			p.MessageType()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(265)
			p.EnumType()
		}

	}

	return localctx
}

// IReservedContext is an interface to support dynamic dispatch.
type IReservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedContext differentiates from other interfaces.
	IsReservedContext()
}

type ReservedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedContext() *ReservedContext {
	var p = new(ReservedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_reserved
	return p
}

func (*ReservedContext) IsReservedContext() {}

func NewReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedContext {
	var p = new(ReservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_reserved

	return p
}

func (s *ReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(IonizerParserRESERVED, 0)
}

func (s *ReservedContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *ReservedContext) Ranges() IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *ReservedContext) ReservedFieldNames() IReservedFieldNamesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedFieldNamesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReservedFieldNamesContext)
}

func (s *ReservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterReserved(s)
	}
}

func (s *ReservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitReserved(s)
	}
}

func (s *ReservedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitReserved(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Reserved() (localctx IReservedContext) {
	localctx = NewReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, IonizerParserRULE_reserved)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(IonizerParserRESERVED)
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IonizerParserINT_LIT:
		{
			p.SetState(269)
			p.Ranges()
		}

	case IonizerParserPROTO3_LIT_SINGLE, IonizerParserPROTO3_LIT_DOBULE, IonizerParserSTR_LIT:
		{
			p.SetState(270)
			p.ReservedFieldNames()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(273)
		p.Match(IonizerParserSEMI)
	}

	return localctx
}

// IExtensionStatementContext is an interface to support dynamic dispatch.
type IExtensionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtensionStatementContext differentiates from other interfaces.
	IsExtensionStatementContext()
}

type ExtensionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionStatementContext() *ExtensionStatementContext {
	var p = new(ExtensionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_extensionStatement
	return p
}

func (*ExtensionStatementContext) IsExtensionStatementContext() {}

func NewExtensionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionStatementContext {
	var p = new(ExtensionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_extensionStatement

	return p
}

func (s *ExtensionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionStatementContext) EXTENSIONS() antlr.TerminalNode {
	return s.GetToken(IonizerParserEXTENSIONS, 0)
}

func (s *ExtensionStatementContext) Ranges() IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *ExtensionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *ExtensionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterExtensionStatement(s)
	}
}

func (s *ExtensionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitExtensionStatement(s)
	}
}

func (s *ExtensionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitExtensionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) ExtensionStatement() (localctx IExtensionStatementContext) {
	localctx = NewExtensionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, IonizerParserRULE_extensionStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Match(IonizerParserEXTENSIONS)
	}
	{
		p.SetState(276)
		p.Ranges()
	}
	{
		p.SetState(277)
		p.Match(IonizerParserSEMI)
	}

	return localctx
}

// IRangesContext is an interface to support dynamic dispatch.
type IRangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangesContext differentiates from other interfaces.
	IsRangesContext()
}

type RangesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesContext() *RangesContext {
	var p = new(RangesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_ranges
	return p
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_ranges

	return p
}

func (s *RangesContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesContext) AllRange_() []IRange_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRange_Context)(nil)).Elem())
	var tst = make([]IRange_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRange_Context)
		}
	}

	return tst
}

func (s *RangesContext) Range_(i int) IRange_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRange_Context)
}

func (s *RangesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserCOMMA)
}

func (s *RangesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserCOMMA, i)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitRanges(s)
	}
}

func (s *RangesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitRanges(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, IonizerParserRULE_ranges)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Range_()
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IonizerParserCOMMA {
		{
			p.SetState(280)
			p.Match(IonizerParserCOMMA)
		}
		{
			p.SetState(281)
			p.Range_()
		}

		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRange_Context is an interface to support dynamic dispatch.
type IRange_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRange_Context differentiates from other interfaces.
	IsRange_Context()
}

type Range_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_Context() *Range_Context {
	var p = new(Range_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_range_
	return p
}

func (*Range_Context) IsRange_Context() {}

func NewRange_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_Context {
	var p = new(Range_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_range_

	return p
}

func (s *Range_Context) GetParser() antlr.Parser { return s.parser }

func (s *Range_Context) AllIntLit() []IIntLitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntLitContext)(nil)).Elem())
	var tst = make([]IIntLitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntLitContext)
		}
	}

	return tst
}

func (s *Range_Context) IntLit(i int) IIntLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntLitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *Range_Context) TO() antlr.TerminalNode {
	return s.GetToken(IonizerParserTO, 0)
}

func (s *Range_Context) MAX() antlr.TerminalNode {
	return s.GetToken(IonizerParserMAX, 0)
}

func (s *Range_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterRange_(s)
	}
}

func (s *Range_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitRange_(s)
	}
}

func (s *Range_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitRange_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Range_() (localctx IRange_Context) {
	localctx = NewRange_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, IonizerParserRULE_range_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.IntLit()
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IonizerParserTO {
		{
			p.SetState(288)
			p.Match(IonizerParserTO)
		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IonizerParserINT_LIT:
			{
				p.SetState(289)
				p.IntLit()
			}

		case IonizerParserMAX:
			{
				p.SetState(290)
				p.Match(IonizerParserMAX)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IReservedFieldNamesContext is an interface to support dynamic dispatch.
type IReservedFieldNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedFieldNamesContext differentiates from other interfaces.
	IsReservedFieldNamesContext()
}

type ReservedFieldNamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedFieldNamesContext() *ReservedFieldNamesContext {
	var p = new(ReservedFieldNamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_reservedFieldNames
	return p
}

func (*ReservedFieldNamesContext) IsReservedFieldNamesContext() {}

func NewReservedFieldNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedFieldNamesContext {
	var p = new(ReservedFieldNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_reservedFieldNames

	return p
}

func (s *ReservedFieldNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedFieldNamesContext) AllStrLit() []IStrLitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStrLitContext)(nil)).Elem())
	var tst = make([]IStrLitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStrLitContext)
		}
	}

	return tst
}

func (s *ReservedFieldNamesContext) StrLit(i int) IStrLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrLitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ReservedFieldNamesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserCOMMA)
}

func (s *ReservedFieldNamesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserCOMMA, i)
}

func (s *ReservedFieldNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedFieldNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedFieldNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterReservedFieldNames(s)
	}
}

func (s *ReservedFieldNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitReservedFieldNames(s)
	}
}

func (s *ReservedFieldNamesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitReservedFieldNames(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) ReservedFieldNames() (localctx IReservedFieldNamesContext) {
	localctx = NewReservedFieldNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, IonizerParserRULE_reservedFieldNames)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.StrLit()
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IonizerParserCOMMA {
		{
			p.SetState(296)
			p.Match(IonizerParserCOMMA)
		}
		{
			p.SetState(297)
			p.StrLit()
		}

		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITopLevelDefContext is an interface to support dynamic dispatch.
type ITopLevelDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopLevelDefContext differentiates from other interfaces.
	IsTopLevelDefContext()
}

type TopLevelDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDefContext() *TopLevelDefContext {
	var p = new(TopLevelDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_topLevelDef
	return p
}

func (*TopLevelDefContext) IsTopLevelDefContext() {}

func NewTopLevelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDefContext {
	var p = new(TopLevelDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_topLevelDef

	return p
}

func (s *TopLevelDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDefContext) MessageDef() IMessageDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageDefContext)
}

func (s *TopLevelDefContext) ExtendDef() IExtendDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtendDefContext)
}

func (s *TopLevelDefContext) EnumDef() IEnumDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *TopLevelDefContext) ServiceDef() IServiceDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceDefContext)
}

func (s *TopLevelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterTopLevelDef(s)
	}
}

func (s *TopLevelDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitTopLevelDef(s)
	}
}

func (s *TopLevelDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitTopLevelDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) TopLevelDef() (localctx ITopLevelDefContext) {
	localctx = NewTopLevelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, IonizerParserRULE_topLevelDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(307)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IonizerParserMESSAGE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(303)
			p.MessageDef()
		}

	case IonizerParserEXTEND:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(304)
			p.ExtendDef()
		}

	case IonizerParserENUM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(305)
			p.EnumDef()
		}

	case IonizerParserSERVICE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(306)
			p.ServiceDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDefContext differentiates from other interfaces.
	IsEnumDefContext()
}

type EnumDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefContext() *EnumDefContext {
	var p = new(EnumDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_enumDef
	return p
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) ENUM() antlr.TerminalNode {
	return s.GetToken(IonizerParserENUM, 0)
}

func (s *EnumDefContext) EnumName() IEnumNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumDefContext) EnumBody() IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitEnumDef(s)
	}
}

func (s *EnumDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitEnumDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, IonizerParserRULE_enumDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(IonizerParserENUM)
	}
	{
		p.SetState(310)
		p.EnumName()
	}
	{
		p.SetState(311)
		p.EnumBody()
	}

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(IonizerParserLC, 0)
}

func (s *EnumBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(IonizerParserRC, 0)
}

func (s *EnumBodyContext) AllEnumElement() []IEnumElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumElementContext)(nil)).Elem())
	var tst = make([]IEnumElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumElementContext)
		}
	}

	return tst
}

func (s *EnumBodyContext) EnumElement(i int) IEnumElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumElementContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (s *EnumBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitEnumBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, IonizerParserRULE_enumBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Match(IonizerParserLC)
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IonizerParserSYNTAX)|(1<<IonizerParserIMPORT)|(1<<IonizerParserWEAK)|(1<<IonizerParserPUBLIC)|(1<<IonizerParserPACKAGE)|(1<<IonizerParserOPTION)|(1<<IonizerParserREPEATED)|(1<<IonizerParserONEOF)|(1<<IonizerParserMAP)|(1<<IonizerParserINT32)|(1<<IonizerParserINT64)|(1<<IonizerParserUINT32)|(1<<IonizerParserUINT64)|(1<<IonizerParserSINT32)|(1<<IonizerParserSINT64)|(1<<IonizerParserFIXED32)|(1<<IonizerParserFIXED64)|(1<<IonizerParserSFIXED32)|(1<<IonizerParserSFIXED64)|(1<<IonizerParserBOOL)|(1<<IonizerParserSTRING)|(1<<IonizerParserDOUBLE)|(1<<IonizerParserFLOAT)|(1<<IonizerParserBYTES)|(1<<IonizerParserRESERVED)|(1<<IonizerParserTO)|(1<<IonizerParserMAX)|(1<<IonizerParserENUM)|(1<<IonizerParserMESSAGE)|(1<<IonizerParserSERVICE)|(1<<IonizerParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IonizerParserEXTENSIONS-32))|(1<<(IonizerParserSTREAM-32))|(1<<(IonizerParserRETURNS-32))|(1<<(IonizerParserEXTEND-32))|(1<<(IonizerParserSEMI-32))|(1<<(IonizerParserBOOL_LIT-32))|(1<<(IonizerParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(314)
			p.EnumElement()
		}

		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(320)
		p.Match(IonizerParserRC)
	}

	return localctx
}

// IEnumElementContext is an interface to support dynamic dispatch.
type IEnumElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumElementContext differentiates from other interfaces.
	IsEnumElementContext()
}

type EnumElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumElementContext() *EnumElementContext {
	var p = new(EnumElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_enumElement
	return p
}

func (*EnumElementContext) IsEnumElementContext() {}

func NewEnumElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumElementContext {
	var p = new(EnumElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_enumElement

	return p
}

func (s *EnumElementContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumElementContext) OptionStatement() IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *EnumElementContext) EnumField() IEnumFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
}

func (s *EnumElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *EnumElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterEnumElement(s)
	}
}

func (s *EnumElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitEnumElement(s)
	}
}

func (s *EnumElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitEnumElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) EnumElement() (localctx IEnumElementContext) {
	localctx = NewEnumElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, IonizerParserRULE_enumElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.OptionStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)
			p.EnumField()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.EmptyStatement_()
		}

	}

	return localctx
}

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_enumField
	return p
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(IonizerParserEQ, 0)
}

func (s *EnumFieldContext) IntLit() IIntLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *EnumFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *EnumFieldContext) MINUS() antlr.TerminalNode {
	return s.GetToken(IonizerParserMINUS, 0)
}

func (s *EnumFieldContext) EnumValueOptions() IEnumValueOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionsContext)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (s *EnumFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitEnumField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, IonizerParserRULE_enumField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Ident()
	}
	{
		p.SetState(328)
		p.Match(IonizerParserEQ)
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IonizerParserMINUS {
		{
			p.SetState(329)
			p.Match(IonizerParserMINUS)
		}

	}
	{
		p.SetState(332)
		p.IntLit()
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IonizerParserLB {
		{
			p.SetState(333)
			p.EnumValueOptions()
		}

	}
	{
		p.SetState(336)
		p.Match(IonizerParserSEMI)
	}

	return localctx
}

// IEnumValueOptionsContext is an interface to support dynamic dispatch.
type IEnumValueOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueOptionsContext differentiates from other interfaces.
	IsEnumValueOptionsContext()
}

type EnumValueOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionsContext() *EnumValueOptionsContext {
	var p = new(EnumValueOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_enumValueOptions
	return p
}

func (*EnumValueOptionsContext) IsEnumValueOptionsContext() {}

func NewEnumValueOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionsContext {
	var p = new(EnumValueOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_enumValueOptions

	return p
}

func (s *EnumValueOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionsContext) LB() antlr.TerminalNode {
	return s.GetToken(IonizerParserLB, 0)
}

func (s *EnumValueOptionsContext) AllEnumValueOption() []IEnumValueOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumValueOptionContext)(nil)).Elem())
	var tst = make([]IEnumValueOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumValueOptionContext)
		}
	}

	return tst
}

func (s *EnumValueOptionsContext) EnumValueOption(i int) IEnumValueOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionContext)
}

func (s *EnumValueOptionsContext) RB() antlr.TerminalNode {
	return s.GetToken(IonizerParserRB, 0)
}

func (s *EnumValueOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserCOMMA)
}

func (s *EnumValueOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserCOMMA, i)
}

func (s *EnumValueOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterEnumValueOptions(s)
	}
}

func (s *EnumValueOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitEnumValueOptions(s)
	}
}

func (s *EnumValueOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitEnumValueOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) EnumValueOptions() (localctx IEnumValueOptionsContext) {
	localctx = NewEnumValueOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, IonizerParserRULE_enumValueOptions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.Match(IonizerParserLB)
	}
	{
		p.SetState(339)
		p.EnumValueOption()
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IonizerParserCOMMA {
		{
			p.SetState(340)
			p.Match(IonizerParserCOMMA)
		}
		{
			p.SetState(341)
			p.EnumValueOption()
		}

		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(347)
		p.Match(IonizerParserRB)
	}

	return localctx
}

// IEnumValueOptionContext is an interface to support dynamic dispatch.
type IEnumValueOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueOptionContext differentiates from other interfaces.
	IsEnumValueOptionContext()
}

type EnumValueOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionContext() *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_enumValueOption
	return p
}

func (*EnumValueOptionContext) IsEnumValueOptionContext() {}

func NewEnumValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_enumValueOption

	return p
}

func (s *EnumValueOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *EnumValueOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(IonizerParserEQ, 0)
}

func (s *EnumValueOptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *EnumValueOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterEnumValueOption(s)
	}
}

func (s *EnumValueOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitEnumValueOption(s)
	}
}

func (s *EnumValueOptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitEnumValueOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) EnumValueOption() (localctx IEnumValueOptionContext) {
	localctx = NewEnumValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, IonizerParserRULE_enumValueOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.OptionName()
	}
	{
		p.SetState(350)
		p.Match(IonizerParserEQ)
	}
	{
		p.SetState(351)
		p.Constant()
	}

	return localctx
}

// IExtendDefContext is an interface to support dynamic dispatch.
type IExtendDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendDefContext differentiates from other interfaces.
	IsExtendDefContext()
}

type ExtendDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendDefContext() *ExtendDefContext {
	var p = new(ExtendDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_extendDef
	return p
}

func (*ExtendDefContext) IsExtendDefContext() {}

func NewExtendDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendDefContext {
	var p = new(ExtendDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_extendDef

	return p
}

func (s *ExtendDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendDefContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(IonizerParserEXTEND, 0)
}

func (s *ExtendDefContext) ExtendName() IExtendNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtendNameContext)
}

func (s *ExtendDefContext) ExtendBody() IExtendBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtendBodyContext)
}

func (s *ExtendDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterExtendDef(s)
	}
}

func (s *ExtendDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitExtendDef(s)
	}
}

func (s *ExtendDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitExtendDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) ExtendDef() (localctx IExtendDefContext) {
	localctx = NewExtendDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, IonizerParserRULE_extendDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Match(IonizerParserEXTEND)
	}
	{
		p.SetState(354)
		p.ExtendName()
	}
	{
		p.SetState(355)
		p.ExtendBody()
	}

	return localctx
}

// IExtendBodyContext is an interface to support dynamic dispatch.
type IExtendBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendBodyContext differentiates from other interfaces.
	IsExtendBodyContext()
}

type ExtendBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendBodyContext() *ExtendBodyContext {
	var p = new(ExtendBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_extendBody
	return p
}

func (*ExtendBodyContext) IsExtendBodyContext() {}

func NewExtendBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendBodyContext {
	var p = new(ExtendBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_extendBody

	return p
}

func (s *ExtendBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(IonizerParserLC, 0)
}

func (s *ExtendBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(IonizerParserRC, 0)
}

func (s *ExtendBodyContext) AllExtendElement() []IExtendElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExtendElementContext)(nil)).Elem())
	var tst = make([]IExtendElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExtendElementContext)
		}
	}

	return tst
}

func (s *ExtendBodyContext) ExtendElement(i int) IExtendElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExtendElementContext)
}

func (s *ExtendBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterExtendBody(s)
	}
}

func (s *ExtendBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitExtendBody(s)
	}
}

func (s *ExtendBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitExtendBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) ExtendBody() (localctx IExtendBodyContext) {
	localctx = NewExtendBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, IonizerParserRULE_extendBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(IonizerParserLC)
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IonizerParserSYNTAX)|(1<<IonizerParserIMPORT)|(1<<IonizerParserWEAK)|(1<<IonizerParserPUBLIC)|(1<<IonizerParserPACKAGE)|(1<<IonizerParserOPTION)|(1<<IonizerParserREPEATED)|(1<<IonizerParserONEOF)|(1<<IonizerParserMAP)|(1<<IonizerParserINT32)|(1<<IonizerParserINT64)|(1<<IonizerParserUINT32)|(1<<IonizerParserUINT64)|(1<<IonizerParserSINT32)|(1<<IonizerParserSINT64)|(1<<IonizerParserFIXED32)|(1<<IonizerParserFIXED64)|(1<<IonizerParserSFIXED32)|(1<<IonizerParserSFIXED64)|(1<<IonizerParserBOOL)|(1<<IonizerParserSTRING)|(1<<IonizerParserDOUBLE)|(1<<IonizerParserFLOAT)|(1<<IonizerParserBYTES)|(1<<IonizerParserRESERVED)|(1<<IonizerParserTO)|(1<<IonizerParserMAX)|(1<<IonizerParserENUM)|(1<<IonizerParserMESSAGE)|(1<<IonizerParserSERVICE)|(1<<IonizerParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IonizerParserEXTENSIONS-32))|(1<<(IonizerParserSTREAM-32))|(1<<(IonizerParserRETURNS-32))|(1<<(IonizerParserEXTEND-32))|(1<<(IonizerParserDOT-32))|(1<<(IonizerParserBOOL_LIT-32))|(1<<(IonizerParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(358)
			p.ExtendElement()
		}

		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(364)
		p.Match(IonizerParserRC)
	}

	return localctx
}

// IExtendElementContext is an interface to support dynamic dispatch.
type IExtendElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendElementContext differentiates from other interfaces.
	IsExtendElementContext()
}

type ExtendElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendElementContext() *ExtendElementContext {
	var p = new(ExtendElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_extendElement
	return p
}

func (*ExtendElementContext) IsExtendElementContext() {}

func NewExtendElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendElementContext {
	var p = new(ExtendElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_extendElement

	return p
}

func (s *ExtendElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendElementContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ExtendElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterExtendElement(s)
	}
}

func (s *ExtendElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitExtendElement(s)
	}
}

func (s *ExtendElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitExtendElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) ExtendElement() (localctx IExtendElementContext) {
	localctx = NewExtendElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, IonizerParserRULE_extendElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Field()
	}

	return localctx
}

// IMessageDefContext is an interface to support dynamic dispatch.
type IMessageDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageDefContext differentiates from other interfaces.
	IsMessageDefContext()
}

type MessageDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageDefContext() *MessageDefContext {
	var p = new(MessageDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_messageDef
	return p
}

func (*MessageDefContext) IsMessageDefContext() {}

func NewMessageDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageDefContext {
	var p = new(MessageDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_messageDef

	return p
}

func (s *MessageDefContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageDefContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(IonizerParserMESSAGE, 0)
}

func (s *MessageDefContext) MessageName() IMessageNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageDefContext) MessageBody() IMessageBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageBodyContext)
}

func (s *MessageDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterMessageDef(s)
	}
}

func (s *MessageDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitMessageDef(s)
	}
}

func (s *MessageDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitMessageDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) MessageDef() (localctx IMessageDefContext) {
	localctx = NewMessageDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, IonizerParserRULE_messageDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(IonizerParserMESSAGE)
	}
	{
		p.SetState(369)
		p.MessageName()
	}
	{
		p.SetState(370)
		p.MessageBody()
	}

	return localctx
}

// IMessageBodyContext is an interface to support dynamic dispatch.
type IMessageBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageBodyContext differentiates from other interfaces.
	IsMessageBodyContext()
}

type MessageBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageBodyContext() *MessageBodyContext {
	var p = new(MessageBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_messageBody
	return p
}

func (*MessageBodyContext) IsMessageBodyContext() {}

func NewMessageBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContext {
	var p = new(MessageBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_messageBody

	return p
}

func (s *MessageBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(IonizerParserLC, 0)
}

func (s *MessageBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(IonizerParserRC, 0)
}

func (s *MessageBodyContext) AllMessageElement() []IMessageElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageElementContext)(nil)).Elem())
	var tst = make([]IMessageElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageElementContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) MessageElement(i int) IMessageElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageElementContext)
}

func (s *MessageBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterMessageBody(s)
	}
}

func (s *MessageBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitMessageBody(s)
	}
}

func (s *MessageBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitMessageBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) MessageBody() (localctx IMessageBodyContext) {
	localctx = NewMessageBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, IonizerParserRULE_messageBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(IonizerParserLC)
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IonizerParserSYNTAX)|(1<<IonizerParserIMPORT)|(1<<IonizerParserWEAK)|(1<<IonizerParserPUBLIC)|(1<<IonizerParserPACKAGE)|(1<<IonizerParserOPTION)|(1<<IonizerParserREPEATED)|(1<<IonizerParserONEOF)|(1<<IonizerParserMAP)|(1<<IonizerParserINT32)|(1<<IonizerParserINT64)|(1<<IonizerParserUINT32)|(1<<IonizerParserUINT64)|(1<<IonizerParserSINT32)|(1<<IonizerParserSINT64)|(1<<IonizerParserFIXED32)|(1<<IonizerParserFIXED64)|(1<<IonizerParserSFIXED32)|(1<<IonizerParserSFIXED64)|(1<<IonizerParserBOOL)|(1<<IonizerParserSTRING)|(1<<IonizerParserDOUBLE)|(1<<IonizerParserFLOAT)|(1<<IonizerParserBYTES)|(1<<IonizerParserRESERVED)|(1<<IonizerParserTO)|(1<<IonizerParserMAX)|(1<<IonizerParserENUM)|(1<<IonizerParserMESSAGE)|(1<<IonizerParserSERVICE)|(1<<IonizerParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IonizerParserEXTENSIONS-32))|(1<<(IonizerParserSTREAM-32))|(1<<(IonizerParserRETURNS-32))|(1<<(IonizerParserEXTEND-32))|(1<<(IonizerParserSEMI-32))|(1<<(IonizerParserDOT-32))|(1<<(IonizerParserBOOL_LIT-32))|(1<<(IonizerParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(373)
			p.MessageElement()
		}

		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(379)
		p.Match(IonizerParserRC)
	}

	return localctx
}

// IMessageElementContext is an interface to support dynamic dispatch.
type IMessageElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageElementContext differentiates from other interfaces.
	IsMessageElementContext()
}

type MessageElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageElementContext() *MessageElementContext {
	var p = new(MessageElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_messageElement
	return p
}

func (*MessageElementContext) IsMessageElementContext() {}

func NewMessageElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageElementContext {
	var p = new(MessageElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_messageElement

	return p
}

func (s *MessageElementContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageElementContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageElementContext) EnumDef() IEnumDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *MessageElementContext) MessageDef() IMessageDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageDefContext)
}

func (s *MessageElementContext) OptionStatement() IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *MessageElementContext) Oneof() IOneofContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOneofContext)
}

func (s *MessageElementContext) MapField() IMapFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapFieldContext)
}

func (s *MessageElementContext) Reserved() IReservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReservedContext)
}

func (s *MessageElementContext) ExtensionStatement() IExtensionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtensionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtensionStatementContext)
}

func (s *MessageElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *MessageElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterMessageElement(s)
	}
}

func (s *MessageElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitMessageElement(s)
	}
}

func (s *MessageElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitMessageElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) MessageElement() (localctx IMessageElementContext) {
	localctx = NewMessageElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, IonizerParserRULE_messageElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(381)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(382)
			p.EnumDef()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(383)
			p.MessageDef()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(384)
			p.OptionStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(385)
			p.Oneof()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(386)
			p.MapField()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(387)
			p.Reserved()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(388)
			p.ExtensionStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(389)
			p.EmptyStatement_()
		}

	}

	return localctx
}

// IServiceDefContext is an interface to support dynamic dispatch.
type IServiceDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceDefContext differentiates from other interfaces.
	IsServiceDefContext()
}

type ServiceDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceDefContext() *ServiceDefContext {
	var p = new(ServiceDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_serviceDef
	return p
}

func (*ServiceDefContext) IsServiceDefContext() {}

func NewServiceDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceDefContext {
	var p = new(ServiceDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_serviceDef

	return p
}

func (s *ServiceDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceDefContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(IonizerParserSERVICE, 0)
}

func (s *ServiceDefContext) ServiceName() IServiceNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceDefContext) LC() antlr.TerminalNode {
	return s.GetToken(IonizerParserLC, 0)
}

func (s *ServiceDefContext) RC() antlr.TerminalNode {
	return s.GetToken(IonizerParserRC, 0)
}

func (s *ServiceDefContext) AllServiceElement() []IServiceElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IServiceElementContext)(nil)).Elem())
	var tst = make([]IServiceElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IServiceElementContext)
		}
	}

	return tst
}

func (s *ServiceDefContext) ServiceElement(i int) IServiceElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IServiceElementContext)
}

func (s *ServiceDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterServiceDef(s)
	}
}

func (s *ServiceDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitServiceDef(s)
	}
}

func (s *ServiceDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitServiceDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) ServiceDef() (localctx IServiceDefContext) {
	localctx = NewServiceDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, IonizerParserRULE_serviceDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Match(IonizerParserSERVICE)
	}
	{
		p.SetState(393)
		p.ServiceName()
	}
	{
		p.SetState(394)
		p.Match(IonizerParserLC)
	}
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IonizerParserOPTION || _la == IonizerParserRPC || _la == IonizerParserSEMI {
		{
			p.SetState(395)
			p.ServiceElement()
		}

		p.SetState(400)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(401)
		p.Match(IonizerParserRC)
	}

	return localctx
}

// IServiceElementContext is an interface to support dynamic dispatch.
type IServiceElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceElementContext differentiates from other interfaces.
	IsServiceElementContext()
}

type ServiceElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceElementContext() *ServiceElementContext {
	var p = new(ServiceElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_serviceElement
	return p
}

func (*ServiceElementContext) IsServiceElementContext() {}

func NewServiceElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceElementContext {
	var p = new(ServiceElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_serviceElement

	return p
}

func (s *ServiceElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceElementContext) OptionStatement() IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *ServiceElementContext) Rpc() IRpcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcContext)
}

func (s *ServiceElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *ServiceElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterServiceElement(s)
	}
}

func (s *ServiceElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitServiceElement(s)
	}
}

func (s *ServiceElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitServiceElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) ServiceElement() (localctx IServiceElementContext) {
	localctx = NewServiceElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, IonizerParserRULE_serviceElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(406)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IonizerParserOPTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(403)
			p.OptionStatement()
		}

	case IonizerParserRPC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(404)
			p.Rpc()
		}

	case IonizerParserSEMI:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(405)
			p.EmptyStatement_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRpcContext is an interface to support dynamic dispatch.
type IRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcContext differentiates from other interfaces.
	IsRpcContext()
}

type RpcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcContext() *RpcContext {
	var p = new(RpcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_rpc
	return p
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) RPC() antlr.TerminalNode {
	return s.GetToken(IonizerParserRPC, 0)
}

func (s *RpcContext) RpcName() IRpcNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcNameContext)
}

func (s *RpcContext) AllLP() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserLP)
}

func (s *RpcContext) LP(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserLP, i)
}

func (s *RpcContext) AllMessageType() []IMessageTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem())
	var tst = make([]IMessageTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageTypeContext)
		}
	}

	return tst
}

func (s *RpcContext) MessageType(i int) IMessageTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *RpcContext) AllRP() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserRP)
}

func (s *RpcContext) RP(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserRP, i)
}

func (s *RpcContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(IonizerParserRETURNS, 0)
}

func (s *RpcContext) LC() antlr.TerminalNode {
	return s.GetToken(IonizerParserLC, 0)
}

func (s *RpcContext) RC() antlr.TerminalNode {
	return s.GetToken(IonizerParserRC, 0)
}

func (s *RpcContext) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *RpcContext) AllSTREAM() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserSTREAM)
}

func (s *RpcContext) STREAM(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserSTREAM, i)
}

func (s *RpcContext) AllOptionStatement() []IOptionStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem())
	var tst = make([]IOptionStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionStatementContext)
		}
	}

	return tst
}

func (s *RpcContext) OptionStatement(i int) IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *RpcContext) AllEmptyStatement_() []IEmptyStatement_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem())
	var tst = make([]IEmptyStatement_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatement_Context)
		}
	}

	return tst
}

func (s *RpcContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *RpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterRpc(s)
	}
}

func (s *RpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitRpc(s)
	}
}

func (s *RpcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitRpc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Rpc() (localctx IRpcContext) {
	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, IonizerParserRULE_rpc)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(408)
		p.Match(IonizerParserRPC)
	}
	{
		p.SetState(409)
		p.RpcName()
	}
	{
		p.SetState(410)
		p.Match(IonizerParserLP)
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(411)
			p.Match(IonizerParserSTREAM)
		}

	}
	{
		p.SetState(414)
		p.MessageType()
	}
	{
		p.SetState(415)
		p.Match(IonizerParserRP)
	}
	{
		p.SetState(416)
		p.Match(IonizerParserRETURNS)
	}
	{
		p.SetState(417)
		p.Match(IonizerParserLP)
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(418)
			p.Match(IonizerParserSTREAM)
		}

	}
	{
		p.SetState(421)
		p.MessageType()
	}
	{
		p.SetState(422)
		p.Match(IonizerParserRP)
	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IonizerParserLC:
		{
			p.SetState(423)
			p.Match(IonizerParserLC)
		}
		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IonizerParserOPTION || _la == IonizerParserSEMI {
			p.SetState(426)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case IonizerParserOPTION:
				{
					p.SetState(424)
					p.OptionStatement()
				}

			case IonizerParserSEMI:
				{
					p.SetState(425)
					p.EmptyStatement_()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(430)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(431)
			p.Match(IonizerParserRC)
		}

	case IonizerParserSEMI:
		{
			p.SetState(432)
			p.Match(IonizerParserSEMI)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ConstantContext) IntLit() IIntLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *ConstantContext) MINUS() antlr.TerminalNode {
	return s.GetToken(IonizerParserMINUS, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(IonizerParserPLUS, 0)
}

func (s *ConstantContext) FloatLit() IFloatLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatLitContext)
}

func (s *ConstantContext) StrLit() IStrLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ConstantContext) BoolLit() IBoolLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolLitContext)
}

func (s *ConstantContext) BlockLit() IBlockLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockLitContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, IonizerParserRULE_constant)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IonizerParserPLUS || _la == IonizerParserMINUS {
			{
				p.SetState(436)
				_la = p.GetTokenStream().LA(1)

				if !(_la == IonizerParserPLUS || _la == IonizerParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(439)
			p.IntLit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IonizerParserPLUS || _la == IonizerParserMINUS {
			{
				p.SetState(440)
				_la = p.GetTokenStream().LA(1)

				if !(_la == IonizerParserPLUS || _la == IonizerParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(443)
			p.FloatLit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(444)
			p.StrLit()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(445)
			p.BoolLit()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(446)
			p.BlockLit()
		}

	}

	return localctx
}

// IBlockLitContext is an interface to support dynamic dispatch.
type IBlockLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockLitContext differentiates from other interfaces.
	IsBlockLitContext()
}

type BlockLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockLitContext() *BlockLitContext {
	var p = new(BlockLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_blockLit
	return p
}

func (*BlockLitContext) IsBlockLitContext() {}

func NewBlockLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockLitContext {
	var p = new(BlockLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_blockLit

	return p
}

func (s *BlockLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockLitContext) LC() antlr.TerminalNode {
	return s.GetToken(IonizerParserLC, 0)
}

func (s *BlockLitContext) RC() antlr.TerminalNode {
	return s.GetToken(IonizerParserRC, 0)
}

func (s *BlockLitContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *BlockLitContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *BlockLitContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserCOLON)
}

func (s *BlockLitContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserCOLON, i)
}

func (s *BlockLitContext) AllConstant() []IConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantContext)(nil)).Elem())
	var tst = make([]IConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantContext)
		}
	}

	return tst
}

func (s *BlockLitContext) Constant(i int) IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *BlockLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterBlockLit(s)
	}
}

func (s *BlockLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitBlockLit(s)
	}
}

func (s *BlockLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitBlockLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) BlockLit() (localctx IBlockLitContext) {
	localctx = NewBlockLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, IonizerParserRULE_blockLit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(449)
		p.Match(IonizerParserLC)
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IonizerParserSYNTAX)|(1<<IonizerParserIMPORT)|(1<<IonizerParserWEAK)|(1<<IonizerParserPUBLIC)|(1<<IonizerParserPACKAGE)|(1<<IonizerParserOPTION)|(1<<IonizerParserREPEATED)|(1<<IonizerParserONEOF)|(1<<IonizerParserMAP)|(1<<IonizerParserINT32)|(1<<IonizerParserINT64)|(1<<IonizerParserUINT32)|(1<<IonizerParserUINT64)|(1<<IonizerParserSINT32)|(1<<IonizerParserSINT64)|(1<<IonizerParserFIXED32)|(1<<IonizerParserFIXED64)|(1<<IonizerParserSFIXED32)|(1<<IonizerParserSFIXED64)|(1<<IonizerParserBOOL)|(1<<IonizerParserSTRING)|(1<<IonizerParserDOUBLE)|(1<<IonizerParserFLOAT)|(1<<IonizerParserBYTES)|(1<<IonizerParserRESERVED)|(1<<IonizerParserTO)|(1<<IonizerParserMAX)|(1<<IonizerParserENUM)|(1<<IonizerParserMESSAGE)|(1<<IonizerParserSERVICE)|(1<<IonizerParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IonizerParserEXTENSIONS-32))|(1<<(IonizerParserSTREAM-32))|(1<<(IonizerParserRETURNS-32))|(1<<(IonizerParserEXTEND-32))|(1<<(IonizerParserBOOL_LIT-32))|(1<<(IonizerParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(450)
			p.Ident()
		}
		{
			p.SetState(451)
			p.Match(IonizerParserCOLON)
		}
		{
			p.SetState(452)
			p.Constant()
		}

		p.SetState(458)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(459)
		p.Match(IonizerParserRC)
	}

	return localctx
}

// IEmptyStatement_Context is an interface to support dynamic dispatch.
type IEmptyStatement_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyStatement_Context differentiates from other interfaces.
	IsEmptyStatement_Context()
}

type EmptyStatement_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatement_Context() *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_emptyStatement_
	return p
}

func (*EmptyStatement_Context) IsEmptyStatement_Context() {}

func NewEmptyStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_emptyStatement_

	return p
}

func (s *EmptyStatement_Context) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatement_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(IonizerParserSEMI, 0)
}

func (s *EmptyStatement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatement_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterEmptyStatement_(s)
	}
}

func (s *EmptyStatement_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitEmptyStatement_(s)
	}
}

func (s *EmptyStatement_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitEmptyStatement_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) EmptyStatement_() (localctx IEmptyStatement_Context) {
	localctx = NewEmptyStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, IonizerParserRULE_emptyStatement_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(461)
		p.Match(IonizerParserSEMI)
	}

	return localctx
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_ident
	return p
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IonizerParserIDENTIFIER, 0)
}

func (s *IdentContext) Keywords() IKeywordsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordsContext)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, IonizerParserRULE_ident)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(465)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IonizerParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(463)
			p.Match(IonizerParserIDENTIFIER)
		}

	case IonizerParserSYNTAX, IonizerParserIMPORT, IonizerParserWEAK, IonizerParserPUBLIC, IonizerParserPACKAGE, IonizerParserOPTION, IonizerParserREPEATED, IonizerParserONEOF, IonizerParserMAP, IonizerParserINT32, IonizerParserINT64, IonizerParserUINT32, IonizerParserUINT64, IonizerParserSINT32, IonizerParserSINT64, IonizerParserFIXED32, IonizerParserFIXED64, IonizerParserSFIXED32, IonizerParserSFIXED64, IonizerParserBOOL, IonizerParserSTRING, IonizerParserDOUBLE, IonizerParserFLOAT, IonizerParserBYTES, IonizerParserRESERVED, IonizerParserTO, IonizerParserMAX, IonizerParserENUM, IonizerParserMESSAGE, IonizerParserSERVICE, IonizerParserRPC, IonizerParserEXTENSIONS, IonizerParserSTREAM, IonizerParserRETURNS, IonizerParserEXTEND, IonizerParserBOOL_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(464)
			p.Keywords()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_fullIdent
	return p
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *FullIdentContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FullIdentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserDOT)
}

func (s *FullIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserDOT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterFullIdent(s)
	}
}

func (s *FullIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitFullIdent(s)
	}
}

func (s *FullIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitFullIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) FullIdent() (localctx IFullIdentContext) {
	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, IonizerParserRULE_fullIdent)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(467)
		p.Ident()
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IonizerParserDOT {
		{
			p.SetState(468)
			p.Match(IonizerParserDOT)
		}
		{
			p.SetState(469)
			p.Ident()
		}

		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMessageNameContext is an interface to support dynamic dispatch.
type IMessageNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageNameContext differentiates from other interfaces.
	IsMessageNameContext()
}

type MessageNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageNameContext() *MessageNameContext {
	var p = new(MessageNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_messageName
	return p
}

func (*MessageNameContext) IsMessageNameContext() {}

func NewMessageNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageNameContext {
	var p = new(MessageNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_messageName

	return p
}

func (s *MessageNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MessageNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterMessageName(s)
	}
}

func (s *MessageNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitMessageName(s)
	}
}

func (s *MessageNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitMessageName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) MessageName() (localctx IMessageNameContext) {
	localctx = NewMessageNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, IonizerParserRULE_messageName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		p.Ident()
	}

	return localctx
}

// IExtendNameContext is an interface to support dynamic dispatch.
type IExtendNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendNameContext differentiates from other interfaces.
	IsExtendNameContext()
}

type ExtendNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendNameContext() *ExtendNameContext {
	var p = new(ExtendNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_extendName
	return p
}

func (*ExtendNameContext) IsExtendNameContext() {}

func NewExtendNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendNameContext {
	var p = new(ExtendNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_extendName

	return p
}

func (s *ExtendNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendNameContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *ExtendNameContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ExtendNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserDOT)
}

func (s *ExtendNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserDOT, i)
}

func (s *ExtendNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterExtendName(s)
	}
}

func (s *ExtendNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitExtendName(s)
	}
}

func (s *ExtendNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitExtendName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) ExtendName() (localctx IExtendNameContext) {
	localctx = NewExtendNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, IonizerParserRULE_extendName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IonizerParserDOT {
		{
			p.SetState(477)
			p.Match(IonizerParserDOT)
		}

	}
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(480)
				p.Ident()
			}
			{
				p.SetState(481)
				p.Match(IonizerParserDOT)
			}

		}
		p.SetState(487)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}
	{
		p.SetState(488)
		p.Ident()
	}

	return localctx
}

// IEnumNameContext is an interface to support dynamic dispatch.
type IEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumNameContext differentiates from other interfaces.
	IsEnumNameContext()
}

type EnumNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumNameContext() *EnumNameContext {
	var p = new(EnumNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_enumName
	return p
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_enumName

	return p
}

func (s *EnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterEnumName(s)
	}
}

func (s *EnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitEnumName(s)
	}
}

func (s *EnumNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitEnumName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) EnumName() (localctx IEnumNameContext) {
	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, IonizerParserRULE_enumName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.Ident()
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (s *FieldNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitFieldName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, IonizerParserRULE_fieldName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.Ident()
	}

	return localctx
}

// IOneofNameContext is an interface to support dynamic dispatch.
type IOneofNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofNameContext differentiates from other interfaces.
	IsOneofNameContext()
}

type OneofNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofNameContext() *OneofNameContext {
	var p = new(OneofNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_oneofName
	return p
}

func (*OneofNameContext) IsOneofNameContext() {}

func NewOneofNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofNameContext {
	var p = new(OneofNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_oneofName

	return p
}

func (s *OneofNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *OneofNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterOneofName(s)
	}
}

func (s *OneofNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitOneofName(s)
	}
}

func (s *OneofNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitOneofName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) OneofName() (localctx IOneofNameContext) {
	localctx = NewOneofNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, IonizerParserRULE_oneofName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(494)
		p.Ident()
	}

	return localctx
}

// IMapNameContext is an interface to support dynamic dispatch.
type IMapNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapNameContext differentiates from other interfaces.
	IsMapNameContext()
}

type MapNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapNameContext() *MapNameContext {
	var p = new(MapNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_mapName
	return p
}

func (*MapNameContext) IsMapNameContext() {}

func NewMapNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapNameContext {
	var p = new(MapNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_mapName

	return p
}

func (s *MapNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MapNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MapNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterMapName(s)
	}
}

func (s *MapNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitMapName(s)
	}
}

func (s *MapNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitMapName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) MapName() (localctx IMapNameContext) {
	localctx = NewMapNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, IonizerParserRULE_mapName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(496)
		p.Ident()
	}

	return localctx
}

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_serviceName
	return p
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterServiceName(s)
	}
}

func (s *ServiceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitServiceName(s)
	}
}

func (s *ServiceNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitServiceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) ServiceName() (localctx IServiceNameContext) {
	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, IonizerParserRULE_serviceName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(498)
		p.Ident()
	}

	return localctx
}

// IRpcNameContext is an interface to support dynamic dispatch.
type IRpcNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcNameContext differentiates from other interfaces.
	IsRpcNameContext()
}

type RpcNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcNameContext() *RpcNameContext {
	var p = new(RpcNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_rpcName
	return p
}

func (*RpcNameContext) IsRpcNameContext() {}

func NewRpcNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcNameContext {
	var p = new(RpcNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_rpcName

	return p
}

func (s *RpcNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *RpcNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterRpcName(s)
	}
}

func (s *RpcNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitRpcName(s)
	}
}

func (s *RpcNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitRpcName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) RpcName() (localctx IRpcNameContext) {
	localctx = NewRpcNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, IonizerParserRULE_rpcName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)
		p.Ident()
	}

	return localctx
}

// IMessageTypeContext is an interface to support dynamic dispatch.
type IMessageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageTypeContext differentiates from other interfaces.
	IsMessageTypeContext()
}

type MessageTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageTypeContext() *MessageTypeContext {
	var p = new(MessageTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_messageType
	return p
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) MessageName() IMessageNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserDOT)
}

func (s *MessageTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserDOT, i)
}

func (s *MessageTypeContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *MessageTypeContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterMessageType(s)
	}
}

func (s *MessageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitMessageType(s)
	}
}

func (s *MessageTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitMessageType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) MessageType() (localctx IMessageTypeContext) {
	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, IonizerParserRULE_messageType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IonizerParserDOT {
		{
			p.SetState(502)
			p.Match(IonizerParserDOT)
		}

	}
	p.SetState(510)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(505)
				p.Ident()
			}
			{
				p.SetState(506)
				p.Match(IonizerParserDOT)
			}

		}
		p.SetState(512)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}
	{
		p.SetState(513)
		p.MessageName()
	}

	return localctx
}

// IEnumTypeContext is an interface to support dynamic dispatch.
type IEnumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeContext differentiates from other interfaces.
	IsEnumTypeContext()
}

type EnumTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeContext() *EnumTypeContext {
	var p = new(EnumTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_enumType
	return p
}

func (*EnumTypeContext) IsEnumTypeContext() {}

func NewEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeContext {
	var p = new(EnumTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_enumType

	return p
}

func (s *EnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeContext) EnumName() IEnumNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(IonizerParserDOT)
}

func (s *EnumTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(IonizerParserDOT, i)
}

func (s *EnumTypeContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *EnumTypeContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterEnumType(s)
	}
}

func (s *EnumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitEnumType(s)
	}
}

func (s *EnumTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitEnumType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) EnumType() (localctx IEnumTypeContext) {
	localctx = NewEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, IonizerParserRULE_enumType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IonizerParserDOT {
		{
			p.SetState(515)
			p.Match(IonizerParserDOT)
		}

	}
	p.SetState(523)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(518)
				p.Ident()
			}
			{
				p.SetState(519)
				p.Match(IonizerParserDOT)
			}

		}
		p.SetState(525)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
	}
	{
		p.SetState(526)
		p.EnumName()
	}

	return localctx
}

// IIntLitContext is an interface to support dynamic dispatch.
type IIntLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntLitContext differentiates from other interfaces.
	IsIntLitContext()
}

type IntLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntLitContext() *IntLitContext {
	var p = new(IntLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_intLit
	return p
}

func (*IntLitContext) IsIntLitContext() {}

func NewIntLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLitContext {
	var p = new(IntLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_intLit

	return p
}

func (s *IntLitContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(IonizerParserINT_LIT, 0)
}

func (s *IntLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterIntLit(s)
	}
}

func (s *IntLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitIntLit(s)
	}
}

func (s *IntLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitIntLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) IntLit() (localctx IIntLitContext) {
	localctx = NewIntLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, IonizerParserRULE_intLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(528)
		p.Match(IonizerParserINT_LIT)
	}

	return localctx
}

// IStrLitContext is an interface to support dynamic dispatch.
type IStrLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrLitContext differentiates from other interfaces.
	IsStrLitContext()
}

type StrLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrLitContext() *StrLitContext {
	var p = new(StrLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_strLit
	return p
}

func (*StrLitContext) IsStrLitContext() {}

func NewStrLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrLitContext {
	var p = new(StrLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_strLit

	return p
}

func (s *StrLitContext) GetParser() antlr.Parser { return s.parser }

func (s *StrLitContext) STR_LIT() antlr.TerminalNode {
	return s.GetToken(IonizerParserSTR_LIT, 0)
}

func (s *StrLitContext) PROTO3_LIT_SINGLE() antlr.TerminalNode {
	return s.GetToken(IonizerParserPROTO3_LIT_SINGLE, 0)
}

func (s *StrLitContext) PROTO3_LIT_DOBULE() antlr.TerminalNode {
	return s.GetToken(IonizerParserPROTO3_LIT_DOBULE, 0)
}

func (s *StrLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterStrLit(s)
	}
}

func (s *StrLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitStrLit(s)
	}
}

func (s *StrLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitStrLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) StrLit() (localctx IStrLitContext) {
	localctx = NewStrLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, IonizerParserRULE_strLit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(530)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(IonizerParserPROTO3_LIT_SINGLE-36))|(1<<(IonizerParserPROTO3_LIT_DOBULE-36))|(1<<(IonizerParserSTR_LIT-36)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBoolLitContext is an interface to support dynamic dispatch.
type IBoolLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolLitContext differentiates from other interfaces.
	IsBoolLitContext()
}

type BoolLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolLitContext() *BoolLitContext {
	var p = new(BoolLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_boolLit
	return p
}

func (*BoolLitContext) IsBoolLitContext() {}

func NewBoolLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLitContext {
	var p = new(BoolLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_boolLit

	return p
}

func (s *BoolLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLitContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(IonizerParserBOOL_LIT, 0)
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterBoolLit(s)
	}
}

func (s *BoolLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitBoolLit(s)
	}
}

func (s *BoolLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitBoolLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) BoolLit() (localctx IBoolLitContext) {
	localctx = NewBoolLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, IonizerParserRULE_boolLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(532)
		p.Match(IonizerParserBOOL_LIT)
	}

	return localctx
}

// IFloatLitContext is an interface to support dynamic dispatch.
type IFloatLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLitContext differentiates from other interfaces.
	IsFloatLitContext()
}

type FloatLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLitContext() *FloatLitContext {
	var p = new(FloatLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_floatLit
	return p
}

func (*FloatLitContext) IsFloatLitContext() {}

func NewFloatLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLitContext {
	var p = new(FloatLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_floatLit

	return p
}

func (s *FloatLitContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(IonizerParserFLOAT_LIT, 0)
}

func (s *FloatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterFloatLit(s)
	}
}

func (s *FloatLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitFloatLit(s)
	}
}

func (s *FloatLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitFloatLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) FloatLit() (localctx IFloatLitContext) {
	localctx = NewFloatLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, IonizerParserRULE_floatLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.Match(IonizerParserFLOAT_LIT)
	}

	return localctx
}

// IKeywordsContext is an interface to support dynamic dispatch.
type IKeywordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordsContext differentiates from other interfaces.
	IsKeywordsContext()
}

type KeywordsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordsContext() *KeywordsContext {
	var p = new(KeywordsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IonizerParserRULE_keywords
	return p
}

func (*KeywordsContext) IsKeywordsContext() {}

func NewKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordsContext {
	var p = new(KeywordsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IonizerParserRULE_keywords

	return p
}

func (s *KeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordsContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(IonizerParserSYNTAX, 0)
}

func (s *KeywordsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(IonizerParserIMPORT, 0)
}

func (s *KeywordsContext) WEAK() antlr.TerminalNode {
	return s.GetToken(IonizerParserWEAK, 0)
}

func (s *KeywordsContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(IonizerParserPUBLIC, 0)
}

func (s *KeywordsContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(IonizerParserPACKAGE, 0)
}

func (s *KeywordsContext) OPTION() antlr.TerminalNode {
	return s.GetToken(IonizerParserOPTION, 0)
}

func (s *KeywordsContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(IonizerParserREPEATED, 0)
}

func (s *KeywordsContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(IonizerParserONEOF, 0)
}

func (s *KeywordsContext) MAP() antlr.TerminalNode {
	return s.GetToken(IonizerParserMAP, 0)
}

func (s *KeywordsContext) INT32() antlr.TerminalNode {
	return s.GetToken(IonizerParserINT32, 0)
}

func (s *KeywordsContext) INT64() antlr.TerminalNode {
	return s.GetToken(IonizerParserINT64, 0)
}

func (s *KeywordsContext) UINT32() antlr.TerminalNode {
	return s.GetToken(IonizerParserUINT32, 0)
}

func (s *KeywordsContext) UINT64() antlr.TerminalNode {
	return s.GetToken(IonizerParserUINT64, 0)
}

func (s *KeywordsContext) SINT32() antlr.TerminalNode {
	return s.GetToken(IonizerParserSINT32, 0)
}

func (s *KeywordsContext) SINT64() antlr.TerminalNode {
	return s.GetToken(IonizerParserSINT64, 0)
}

func (s *KeywordsContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(IonizerParserFIXED32, 0)
}

func (s *KeywordsContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(IonizerParserFIXED64, 0)
}

func (s *KeywordsContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(IonizerParserSFIXED32, 0)
}

func (s *KeywordsContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(IonizerParserSFIXED64, 0)
}

func (s *KeywordsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(IonizerParserBOOL, 0)
}

func (s *KeywordsContext) STRING() antlr.TerminalNode {
	return s.GetToken(IonizerParserSTRING, 0)
}

func (s *KeywordsContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(IonizerParserDOUBLE, 0)
}

func (s *KeywordsContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(IonizerParserFLOAT, 0)
}

func (s *KeywordsContext) BYTES() antlr.TerminalNode {
	return s.GetToken(IonizerParserBYTES, 0)
}

func (s *KeywordsContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(IonizerParserRESERVED, 0)
}

func (s *KeywordsContext) TO() antlr.TerminalNode {
	return s.GetToken(IonizerParserTO, 0)
}

func (s *KeywordsContext) MAX() antlr.TerminalNode {
	return s.GetToken(IonizerParserMAX, 0)
}

func (s *KeywordsContext) ENUM() antlr.TerminalNode {
	return s.GetToken(IonizerParserENUM, 0)
}

func (s *KeywordsContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(IonizerParserMESSAGE, 0)
}

func (s *KeywordsContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(IonizerParserSERVICE, 0)
}

func (s *KeywordsContext) EXTENSIONS() antlr.TerminalNode {
	return s.GetToken(IonizerParserEXTENSIONS, 0)
}

func (s *KeywordsContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(IonizerParserEXTEND, 0)
}

func (s *KeywordsContext) RPC() antlr.TerminalNode {
	return s.GetToken(IonizerParserRPC, 0)
}

func (s *KeywordsContext) STREAM() antlr.TerminalNode {
	return s.GetToken(IonizerParserSTREAM, 0)
}

func (s *KeywordsContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(IonizerParserRETURNS, 0)
}

func (s *KeywordsContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(IonizerParserBOOL_LIT, 0)
}

func (s *KeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.EnterKeywords(s)
	}
}

func (s *KeywordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IonizerListener); ok {
		listenerT.ExitKeywords(s)
	}
}

func (s *KeywordsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IonizerVisitor:
		return t.VisitKeywords(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IonizerParser) Keywords() (localctx IKeywordsContext) {
	localctx = NewKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, IonizerParserRULE_keywords)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(536)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IonizerParserSYNTAX)|(1<<IonizerParserIMPORT)|(1<<IonizerParserWEAK)|(1<<IonizerParserPUBLIC)|(1<<IonizerParserPACKAGE)|(1<<IonizerParserOPTION)|(1<<IonizerParserREPEATED)|(1<<IonizerParserONEOF)|(1<<IonizerParserMAP)|(1<<IonizerParserINT32)|(1<<IonizerParserINT64)|(1<<IonizerParserUINT32)|(1<<IonizerParserUINT64)|(1<<IonizerParserSINT32)|(1<<IonizerParserSINT64)|(1<<IonizerParserFIXED32)|(1<<IonizerParserFIXED64)|(1<<IonizerParserSFIXED32)|(1<<IonizerParserSFIXED64)|(1<<IonizerParserBOOL)|(1<<IonizerParserSTRING)|(1<<IonizerParserDOUBLE)|(1<<IonizerParserFLOAT)|(1<<IonizerParserBYTES)|(1<<IonizerParserRESERVED)|(1<<IonizerParserTO)|(1<<IonizerParserMAX)|(1<<IonizerParserENUM)|(1<<IonizerParserMESSAGE)|(1<<IonizerParserSERVICE)|(1<<IonizerParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IonizerParserEXTENSIONS-32))|(1<<(IonizerParserSTREAM-32))|(1<<(IonizerParserRETURNS-32))|(1<<(IonizerParserEXTEND-32))|(1<<(IonizerParserBOOL_LIT-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
