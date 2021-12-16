
[ table_id_object:
00a6823403ea3055  roxe.token (code)
454f530000000000  "ROC" symbol_code (scope)
0000000000904dc6  stat (table name)
00a6823403ea3055  roxe.token (payer?)
01000000          count ?
]

[ size of objects in table:
01                size, as unsigned_int (varint) = 1
]

[ key value object:
454f530000000000  "ROC", primary key?
00a6823403ea3055  roxe.token, payer ?

28                length=40 ?
835fd1f546090000   uint64 = 10200376500099, amount 1,020,037,650.0099 ?!
04454f5300000000  "4,ROC"
00407a10f35a0000   uint64 = 100000000000000, amount 10,000,000,000.0000
04454f5300000000  "4,ROC"
0000000000ea3055   roxe
   `-> these last few fields seem to correspond to: https://roxeq.app/account/roxe.token/tables?lowerBound=&scope=ROC&tableName=stat
]


00000000
00

[ table_id_object:
0000000000ea3055  "roxe"  (code)
0000000000ea3055  "roxe"  (scope)
0000c80a5e23a5b9  "rammarket" name-encoded (table name)
0000000000ea3055  "roxe" (payer)
01000000          count ? or size ? number of elements in table
]

[ size of objects in table:
01                varuint = 1
]

[ key value object:
0452414d434f5245  "4,RAMCORE"  primary_key?
0000000000ea3055  "roxe"  payer?

40                length=64 ? length of shared_blob

00407a10f35a0000  uint64 = 100000000000000, 10,000,000,000.0000
0452414d434f5245  "4,RAMCORE"

458a958919000000  uint64 = 109682461253, 109,682,461,253
0052414d00000000  "0,RAM"
000000000000e03f  float64 = 0.50000

1f8274480d000000  uint64 = 57050169887, 5,705,016.9887
04454f5300000000  "4,ROC"
000000000000e03f  float64 = 0.50000
   `-> looks like: https://roxeq.app/account/roxe/tables?lowerBound=&scope=roxe&tableName=rammarket
]


00000000
00
0000000000ea3055  "roxe" (code)
0000000000ea3055  "roxe" (scope)
0000000044736864  "global" (table name)
0000000000ea3055  "roxe" (payer)
01000000          elements counts
01
0000000044736864  "global" (prim key)
0000000000ea3055  "roxe" (payer)
9a01              blob_size: 282?
0000100000000000
e803000000000800
0c000000f4010000
1400000064000000
400d0300c4090000f049020064000000100e00005802000080533b00

00100000060006000000e4902c000000bb754e07130000001f9e68f40a000000d638e44ca0b5733ba3a70500f516670b0000000058123b010000000096df00001da5259d8a05000000e15fd7136f050015007414edb8a3022344e51de44c00000000000000000000ea3055000000000000403800000000ab7b15d60000000000004038010000000100000000000040380000000000004038300000000000004038c8d52aef6900000004454f5300000000f39710995d00000004454f53000000001705e6fc070000000000000000
00a6823403ea3055000090e602ea3055000000384f4d11320000000000ea30550100000001454f530000000000000090e602ea3055108e166ef40a00000004454f5300000000000000000000a6823403ea3055a0d492e602ea3055000000384f4d11320000000000ea30550100000001454f530000000000a0d492e602ea305510918658910400000004454f530000000000000000000000000000ea30550000000000004038000000204d73a24a00000000000040380500000005000000000000403800000000000040383000000000000040380000000000004038b8ae2aef6900000004454f5300000000e37010995d00000004454f5300000000004038205c3599c30000000000004038300000000000004038004038205c3599c380841e000000000004454f530000000080841e000000000004454f5300000000004038405c3599c30000000000004038300000000000004038004038405c3599c380841e000000000004454f530000000080841e000000000004454f5300000000004038605c3599c30000000000004038300000000000004038004038605c3599c380841e000000000004454f530000000080841e000000000004454f5300000000000000201c108ec60000000000004038300000000000004038000000201c108ec680841e000000000004454f530000000080