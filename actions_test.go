package roxe

import (
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActionUnmarshalBinary(t *testing.T) {
	tests := []struct {
		in     string
		jsonTx string
	}{
		{
			"0000000000ea30550000e62e0061736d0100000001621260037f7e7f0060057f7e7e7e7e0060047f7e7e7e0060027f7f0060027f7e006000006000017e60027e7e006000017f60027f7f017f60037f7f7f017f60017e0060027f7f017e60047e7e7e7e0060027e7f0060017f0060037e7e7e0060017f017f02ee010c03656e760561626f7274000503656e7610616374696f6e5f646174615f73697a65000803656e760c63757272656e745f74696d65000603656e760c656f73696f5f617373657274000303656e76066d656d637079000a03656e76066d656d736574000a03656e7610726561645f616374696f6e5f64617461000903656e760c726571756972655f61757468000b03656e760d726571756972655f6175746832000703656e760e7365745f70726976696c65676564000e03656e76167365745f70726f706f7365645f70726f647563657273000c03656e76137365745f7265736f757263655f6c696d697473000d031d1c090909080f1000090109020903090409090303110f0f0a1109110f0504050170010606050301000107af010a066d656d6f72790200165f5a6571524b3131636865636b73756d32353653315f000c165f5a6571524b3131636865636b73756d31363053315f000d165f5a6e65524b3131636865636b73756d31363053315f000e036e6f77000f305f5a4e35656f73696f3132726571756972655f6175746845524b4e535f31367065726d697373696f6e5f6c6576656c450010056170706c790011066d656d636d700022066d616c6c6f63002304667265650026090c010041000b0627121416181a0abf281c0b002000200141201022450b0b002000200141201022450b0d0020002001412010224100470b0a00100242c0843d80a70b0e002000290300200029030810080bf60603027f047e017f4100410028020441e0006b220936020442002106423b2105411021044200210703400240024002400240024020064206560d0020042c00002203419f7f6a41ff017141194b0d01200341a5016a21030c020b420021082006420b580d020c030b200341d0016a41002003414f6a41ff01714105491b21030b2003ad42388642388721080b2008421f83200542ffffffff0f838621080b200441016a2104200642017c2106200820078421072005427b7c2205427a520d000b024020072002520d0042002106423b2105412021044200210703400240024002400240024020064204560d0020042c00002203419f7f6a41ff017141194b0d01200341a5016a21030c020b420021082006420b580d020c030b200341d0016a41002003414f6a41ff01714105491b21030b2003ad42388642388721080b2008421f83200542ffffffff0f838621080b200441016a2104200642017c2106200820078421072005427b7c2205427a520d000b2007200151413010030b0240024020012000510d0042002106423b2105411021044200210703400240024002400240024020064206560d0020042c00002203419f7f6a41ff017141194b0d01200341a5016a21030c020b420021082006420b580d020c030b200341d0016a41002003414f6a41ff01714105491b21030b2003ad42388642388721080b2008421f83200542ffffffff0f838621080b200441016a2104200642017c2106200820078421072005427b7c2205427a520d000b20072002520d010b200920003703580240024002400240200242ffffb7f6a497b2d942570d002002428080b8f6a497b2d942510d01200242808080c093fad6d942510d0220024280808080b6f7d6d942520d04200941003602542009410136025020092009290350370208200941d8006a200941086a10131a0c040b20024280808080daac9bd6ba7f510d022002428080b8f6a4979ad942520d032009410036024c2009410236024820092009290348370210200941d8006a200941106a10151a0c030b200941003602442009410336024020092009290340370218200941d8006a200941186a10171a0c020b2009410036023c2009410436023820092009290338370220200941d8006a200941206a10191a0c010b200941003602342009410536023020092009290330370228200941d8006a200941286a101b1a0b4100200941e0006a3602040b1200200029030010072001200241004710090bfc0103017f017e047f410028020441106b2207210641002007360204200128020421022001280200210541002104024010012201450d00024002402001418104490d002001102321040c010b410020072001410f6a4170716b22043602040b2004200110061a0b200641003a000820064200370300200141074b41f000100320062004410810041a200141084741f0001003200641086a2207200441086a410110041a02402001418104490d00200410260b200020024101756a210120072d000021042006290300210302402002410171450d00200128020020056a28020021050b20012003200441ff017120051100004100200641106a36020441010b1300200029030010072001200220032004100b0baf0203027f047e037f410028020441c0006b220a21094100200a3602042001280204210220012802002108024002400240024010012203450d002003418104490d012003102321010c020b410021010c020b4100200a2003410f6a4170716b22013602040b2001200310061a0b2009420037030820094200370300200942003703102009420037031820092001360224200920013602202009200120036a3602282009200941206a36023020092009360238200941386a200941306a101e02402003418104490d00200110260b200020024101756a2101200941186a2903002107200941106a2903002106200941086a29030021052009290300210402402002410171450d00200128020020086a28020021080b2001200420052006200720081101004100200941c0006a36020441010b0900200029030010070bb60203017f037e057f410028020441206b2208210a410020083602042001280204210220012802002109024002400240024010012201450d002001418104490d012001102321080c020b410021080c020b410020082001410f6a4170716b22083602040b2008200110061a0b200a4200370310200a4200370308200a4200370318200141074b41f0001003200a41086a2008410810041a2001417871220641084741f0001003200a41086a41086a2207200841086a410810041a200641104741f0001003200a41086a41106a2206200841106a410810041a02402001418104490d00200810260b200020024101756a21012006290300210520072903002104200a290308210302402002410171450d00200128020020096a28020021090b200120032004200520091102004100200a41206a36020441010b5301027f410028020422022103200029030010070240024010012200418104490d002000102321020c010b410020022000410f6a4170716b22023602040b2002200010061a20022000100a1a410020033602040bc20401057f410028020441306b2206210541002006360204200128020421022001280200210441002101024010012203450d00024002402003418104490d002003102321010c010b410020062003410f6a4170716b22013602040b2001200310061a0b200541003602082005420037030020052001360224200520013602202005200120036a360228200541206a2005101c1a02402003418104490d00200110260b200542003703104100210120054100360218200528020420052802006b220341306d21060240024002402003450d00200641d6aad52a4f0d01200541186a2003101f2201200641306c6a36020020052001360210200520013602142005280204200528020022066b22034101480d0020012006200310041a20052005280214200341306e41306c6a22013602140b200020024101756a210302402002410171450d00200328020020046a28020021040b2005420037032020054100360228200120052802106b220141306d210202402001450d00200241d6aad52a4f0d02200541286a2001101f2201200241306c6a36020020052001360220200520013602242005280214200528021022066b22024101480d0020012006200210041a20052001200241306e41306c6a3602240b2003200541206a2004110300024020052802202201450d0020052001360224200110200b024020052802102201450d0020052001360214200110200b024020052802002201450d0020052001360204200110200b4100200541306a36020441010f0b200541106a1021000b200541206a1021000b0600200110070bd50103017f017e037f410028020441106b22042106410020043602042001280204210220012802002105024002400240024010012201450d002001418104490d012001102321040c020b410021040c020b410020042001410f6a4170716b22043602040b2004200110061a0b20064200370308200141074b41f0001003200641086a2004410810041a2006290308210302402001418104490d00200410260b200020024101756a210102402002410171450d00200128020020056a28020021050b2001200320051104004100200641106a36020441010bd20203037f017e027f200028020421074100210642002105200041086a2102200041046a2103034020072002280200494180011003200328020022072d000021042003200741016a2207360200200441ff0071200641ff0171220674ad2005842105200641076a210620044107760d000b0240024002402005a7220420012802042202200128020022076b41306d22064d0d002001200420066b101d20012802002207200141046a2802002202470d010c020b0240200420064f0d00200141046a2007200441306c6a22023602000b20072002460d010b200041046a220428020021060340200041086a220328020020066b41074b41f000100320072004280200410810041a2004200428020041086a2206360200200328020020066b41214b41f0001003200741086a2004280200412210041a2004200428020041226a2206360200200741306a22072002470d000b0b20000bc70201057f0240024002400240024020002802082202200028020422066b41306d20014f0d002006200028020022056b41306d220320016a220441d6aad52a4f0d0241d5aad52a21060240200220056b41306d220241a9d5aa154b0d0020042002410174220620062004491b2206450d020b200641306c101f21020c030b200041046a2100034020064100413010051a2000200028020041306a22063602002001417f6a22010d000c040b0b41002106410021020c010b20001021000b2002200641306c6a21042002200341306c6a220521060340200641004130100541306a21062001417f6a22010d000b2005200041046a2203280200200028020022016b220241506d41306c6a2105024020024101480d0020052001200210041a200028020021010b2000200536020020032006360200200041086a20043602002001450d00200110200f0b0bdf0101027f200028020021032001280200220228020820022802046b41074b41f000100320032002280204410810041a2002200228020441086a360204200028020021002001280200220228020820022802046b41074b41f0001003200041086a2002280204410810041a2002200228020441086a3602042001280200220228020820022802046b41074b41f0001003200041106a2002280204410810041a2002200228020441086a3602042001280200220128020820012802046b41074b41f0001003200041186a2001280204410810041a2001200128020441086a3602040b3801027f02402000410120001b2201102322000d000340410021004100280284012202450d012002110500200110232200450d000b0b20000b0e0002402000450d00200010260b0b05001000000b4901037f4100210502402002450d000240034020002d0000220320012d00002204470d01200141016a2101200041016a21002002417f6a22020d000c020b0b200320046b21050b20050b0900418801200010240bcd04010c7f02402001450d00024020002802c041220d0d004110210d200041c0c1006a41103602000b200141086a200141046a41077122026b200120021b210202400240024020002802c441220a200d4f0d002000200a410c6c6a4180c0006a21010240200a0d0020004184c0006a220d2802000d0020014180c000360200200d20003602000b200241046a210a034002402001280208220d200a6a20012802004b0d002001280204200d6a220d200d28020041808080807871200272360200200141086a22012001280200200a6a360200200d200d28020041808080807872360200200d41046a22010d030b2000102522010d000b0b41fcffffff0720026b2104200041c8c1006a210b200041c0c1006a210c20002802c8412203210d03402000200d410c6c6a22014188c0006a28020020014180c0006a22052802004641e0c200100320014184c0006a280200220641046a210d0340200620052802006a2107200d417c6a2208280200220941ffffffff07712101024020094100480d000240200120024f0d000340200d20016a220a20074f0d01200a280200220a4100480d012001200a41ffffffff07716a41046a22012002490d000b0b20082001200220012002491b200941808080807871723602000240200120024d0d00200d20026a200420016a41ffffffff07713602000b200120024f0d040b200d20016a41046a220d2007490d000b41002101200b4100200b28020041016a220d200d200c280200461b220d360200200d2003470d000b0b20010f0b2008200828020041808080807872360200200d0f0b41000b870501087f20002802c44121010240024041002d00b643450d0041002802b84321070c010b3f002107410041013a00b6434100200741107422073602b8430b200721030240024002400240200741ffff036a41107622023f0022084d0d00200220086b40001a4100210820023f00470d0141002802b84321030b41002108410020033602b84320074100480d0020002001410c6c6a210220074180800441808008200741ffff037122084181f8034922061b6a2008200741ffff077120061b6b20076b2107024041002d00b6430d003f002103410041013a00b6434100200341107422033602b8430b20024180c0006a210220074100480d01200321060240200741076a417871220520036a41ffff036a41107622083f0022044d0d00200820046b40001a20083f00470d0241002802b84321060b4100200620056a3602b8432003417f460d0120002001410c6c6a22014184c0006a2802002206200228020022086a2003460d020240200820014188c0006a22052802002201460d00200620016a2206200628020041808080807871417c20016b20086a72360200200520022802003602002006200628020041ffffffff07713602000b200041c4c1006a2202200228020041016a220236020020002002410c6c6a22004184c0006a200336020020004180c0006a220820073602000b20080f0b02402002280200220820002001410c6c6a22034188c0006a22012802002207460d0020034184c0006a28020020076a2203200328020041808080807871417c20076b20086a72360200200120022802003602002003200328020041ffffffff07713602000b2000200041c4c1006a220728020041016a22033602c0412007200336020041000f0b2002200820076a36020020020b7b01037f024002402000450d0041002802c84222024101480d004188c10021032002410c6c4188c1006a21010340200341046a2802002202450d010240200241046a20004b0d00200220032802006a20004b0d030b2003410c6a22032001490d000b0b0f0b2000417c6a2203200328020041ffffffff07713602000b0300000b0bd901070041040b04c04900000041100b086f6e6572726f72000041200b06656f73696f000041300b406f6e6572726f7220616374696f6e277320617265206f6e6c792076616c69642066726f6d207468652022656f73696f222073797374656d206163636f756e74000041f0000b057265616400004180010b04676574000041e0c2000b566d616c6c6f635f66726f6d5f6672656564207761732064657369676e656420746f206f6e6c792062652063616c6c6564206166746572205f686561702077617320636f6d706c6574656c7920616c6c6f636174656400",
			`{"account":"roxe","vmtype":0,"vmversion":0,"code":"0061736d0100000001621260037f7e7f0060057f7e7e7e7e0060047f7e7e7e0060027f7f0060027f7e006000006000017e60027e7e006000017f60027f7f017f60037f7f7f017f60017e0060027f7f017e60047e7e7e7e0060027e7f0060017f0060037e7e7e0060017f017f02ee010c03656e760561626f7274000503656e7610616374696f6e5f646174615f73697a65000803656e760c63757272656e745f74696d65000603656e760c656f73696f5f617373657274000303656e76066d656d637079000a03656e76066d656d736574000a03656e7610726561645f616374696f6e5f64617461000903656e760c726571756972655f61757468000b03656e760d726571756972655f6175746832000703656e760e7365745f70726976696c65676564000e03656e76167365745f70726f706f7365645f70726f647563657273000c03656e76137365745f7265736f757263655f6c696d697473000d031d1c090909080f1000090109020903090409090303110f0f0a1109110f0504050170010606050301000107af010a066d656d6f72790200165f5a6571524b3131636865636b73756d32353653315f000c165f5a6571524b3131636865636b73756d31363053315f000d165f5a6e65524b3131636865636b73756d31363053315f000e036e6f77000f305f5a4e35656f73696f3132726571756972655f6175746845524b4e535f31367065726d697373696f6e5f6c6576656c450010056170706c790011066d656d636d700022066d616c6c6f63002304667265650026090c010041000b0627121416181a0abf281c0b002000200141201022450b0b002000200141201022450b0d0020002001412010224100470b0a00100242c0843d80a70b0e002000290300200029030810080bf60603027f047e017f4100410028020441e0006b220936020442002106423b2105411021044200210703400240024002400240024020064206560d0020042c00002203419f7f6a41ff017141194b0d01200341a5016a21030c020b420021082006420b580d020c030b200341d0016a41002003414f6a41ff01714105491b21030b2003ad42388642388721080b2008421f83200542ffffffff0f838621080b200441016a2104200642017c2106200820078421072005427b7c2205427a520d000b024020072002520d0042002106423b2105412021044200210703400240024002400240024020064204560d0020042c00002203419f7f6a41ff017141194b0d01200341a5016a21030c020b420021082006420b580d020c030b200341d0016a41002003414f6a41ff01714105491b21030b2003ad42388642388721080b2008421f83200542ffffffff0f838621080b200441016a2104200642017c2106200820078421072005427b7c2205427a520d000b2007200151413010030b0240024020012000510d0042002106423b2105411021044200210703400240024002400240024020064206560d0020042c00002203419f7f6a41ff017141194b0d01200341a5016a21030c020b420021082006420b580d020c030b200341d0016a41002003414f6a41ff01714105491b21030b2003ad42388642388721080b2008421f83200542ffffffff0f838621080b200441016a2104200642017c2106200820078421072005427b7c2205427a520d000b20072002520d010b200920003703580240024002400240200242ffffb7f6a497b2d942570d002002428080b8f6a497b2d942510d01200242808080c093fad6d942510d0220024280808080b6f7d6d942520d04200941003602542009410136025020092009290350370208200941d8006a200941086a10131a0c040b20024280808080daac9bd6ba7f510d022002428080b8f6a4979ad942520d032009410036024c2009410236024820092009290348370210200941d8006a200941106a10151a0c030b200941003602442009410336024020092009290340370218200941d8006a200941186a10171a0c020b2009410036023c2009410436023820092009290338370220200941d8006a200941206a10191a0c010b200941003602342009410536023020092009290330370228200941d8006a200941286a101b1a0b4100200941e0006a3602040b1200200029030010072001200241004710090bfc0103017f017e047f410028020441106b2207210641002007360204200128020421022001280200210541002104024010012201450d00024002402001418104490d002001102321040c010b410020072001410f6a4170716b22043602040b2004200110061a0b200641003a000820064200370300200141074b41f000100320062004410810041a200141084741f0001003200641086a2207200441086a410110041a02402001418104490d00200410260b200020024101756a210120072d000021042006290300210302402002410171450d00200128020020056a28020021050b20012003200441ff017120051100004100200641106a36020441010b1300200029030010072001200220032004100b0baf0203027f047e037f410028020441c0006b220a21094100200a3602042001280204210220012802002108024002400240024010012203450d002003418104490d012003102321010c020b410021010c020b4100200a2003410f6a4170716b22013602040b2001200310061a0b2009420037030820094200370300200942003703102009420037031820092001360224200920013602202009200120036a3602282009200941206a36023020092009360238200941386a200941306a101e02402003418104490d00200110260b200020024101756a2101200941186a2903002107200941106a2903002106200941086a29030021052009290300210402402002410171450d00200128020020086a28020021080b2001200420052006200720081101004100200941c0006a36020441010b0900200029030010070bb60203017f037e057f410028020441206b2208210a410020083602042001280204210220012802002109024002400240024010012201450d002001418104490d012001102321080c020b410021080c020b410020082001410f6a4170716b22083602040b2008200110061a0b200a4200370310200a4200370308200a4200370318200141074b41f0001003200a41086a2008410810041a2001417871220641084741f0001003200a41086a41086a2207200841086a410810041a200641104741f0001003200a41086a41106a2206200841106a410810041a02402001418104490d00200810260b200020024101756a21012006290300210520072903002104200a290308210302402002410171450d00200128020020096a28020021090b200120032004200520091102004100200a41206a36020441010b5301027f410028020422022103200029030010070240024010012200418104490d002000102321020c010b410020022000410f6a4170716b22023602040b2002200010061a20022000100a1a410020033602040bc20401057f410028020441306b2206210541002006360204200128020421022001280200210441002101024010012203450d00024002402003418104490d002003102321010c010b410020062003410f6a4170716b22013602040b2001200310061a0b200541003602082005420037030020052001360224200520013602202005200120036a360228200541206a2005101c1a02402003418104490d00200110260b200542003703104100210120054100360218200528020420052802006b220341306d21060240024002402003450d00200641d6aad52a4f0d01200541186a2003101f2201200641306c6a36020020052001360210200520013602142005280204200528020022066b22034101480d0020012006200310041a20052005280214200341306e41306c6a22013602140b200020024101756a210302402002410171450d00200328020020046a28020021040b2005420037032020054100360228200120052802106b220141306d210202402001450d00200241d6aad52a4f0d02200541286a2001101f2201200241306c6a36020020052001360220200520013602242005280214200528021022066b22024101480d0020012006200210041a20052001200241306e41306c6a3602240b2003200541206a2004110300024020052802202201450d0020052001360224200110200b024020052802102201450d0020052001360214200110200b024020052802002201450d0020052001360204200110200b4100200541306a36020441010f0b200541106a1021000b200541206a1021000b0600200110070bd50103017f017e037f410028020441106b22042106410020043602042001280204210220012802002105024002400240024010012201450d002001418104490d012001102321040c020b410021040c020b410020042001410f6a4170716b22043602040b2004200110061a0b20064200370308200141074b41f0001003200641086a2004410810041a2006290308210302402001418104490d00200410260b200020024101756a210102402002410171450d00200128020020056a28020021050b2001200320051104004100200641106a36020441010bd20203037f017e027f200028020421074100210642002105200041086a2102200041046a2103034020072002280200494180011003200328020022072d000021042003200741016a2207360200200441ff0071200641ff0171220674ad2005842105200641076a210620044107760d000b0240024002402005a7220420012802042202200128020022076b41306d22064d0d002001200420066b101d20012802002207200141046a2802002202470d010c020b0240200420064f0d00200141046a2007200441306c6a22023602000b20072002460d010b200041046a220428020021060340200041086a220328020020066b41074b41f000100320072004280200410810041a2004200428020041086a2206360200200328020020066b41214b41f0001003200741086a2004280200412210041a2004200428020041226a2206360200200741306a22072002470d000b0b20000bc70201057f0240024002400240024020002802082202200028020422066b41306d20014f0d002006200028020022056b41306d220320016a220441d6aad52a4f0d0241d5aad52a21060240200220056b41306d220241a9d5aa154b0d0020042002410174220620062004491b2206450d020b200641306c101f21020c030b200041046a2100034020064100413010051a2000200028020041306a22063602002001417f6a22010d000c040b0b41002106410021020c010b20001021000b2002200641306c6a21042002200341306c6a220521060340200641004130100541306a21062001417f6a22010d000b2005200041046a2203280200200028020022016b220241506d41306c6a2105024020024101480d0020052001200210041a200028020021010b2000200536020020032006360200200041086a20043602002001450d00200110200f0b0bdf0101027f200028020021032001280200220228020820022802046b41074b41f000100320032002280204410810041a2002200228020441086a360204200028020021002001280200220228020820022802046b41074b41f0001003200041086a2002280204410810041a2002200228020441086a3602042001280200220228020820022802046b41074b41f0001003200041106a2002280204410810041a2002200228020441086a3602042001280200220128020820012802046b41074b41f0001003200041186a2001280204410810041a2001200128020441086a3602040b3801027f02402000410120001b2201102322000d000340410021004100280284012202450d012002110500200110232200450d000b0b20000b0e0002402000450d00200010260b0b05001000000b4901037f4100210502402002450d000240034020002d0000220320012d00002204470d01200141016a2101200041016a21002002417f6a22020d000c020b0b200320046b21050b20050b0900418801200010240bcd04010c7f02402001450d00024020002802c041220d0d004110210d200041c0c1006a41103602000b200141086a200141046a41077122026b200120021b210202400240024020002802c441220a200d4f0d002000200a410c6c6a4180c0006a21010240200a0d0020004184c0006a220d2802000d0020014180c000360200200d20003602000b200241046a210a034002402001280208220d200a6a20012802004b0d002001280204200d6a220d200d28020041808080807871200272360200200141086a22012001280200200a6a360200200d200d28020041808080807872360200200d41046a22010d030b2000102522010d000b0b41fcffffff0720026b2104200041c8c1006a210b200041c0c1006a210c20002802c8412203210d03402000200d410c6c6a22014188c0006a28020020014180c0006a22052802004641e0c200100320014184c0006a280200220641046a210d0340200620052802006a2107200d417c6a2208280200220941ffffffff07712101024020094100480d000240200120024f0d000340200d20016a220a20074f0d01200a280200220a4100480d012001200a41ffffffff07716a41046a22012002490d000b0b20082001200220012002491b200941808080807871723602000240200120024d0d00200d20026a200420016a41ffffffff07713602000b200120024f0d040b200d20016a41046a220d2007490d000b41002101200b4100200b28020041016a220d200d200c280200461b220d360200200d2003470d000b0b20010f0b2008200828020041808080807872360200200d0f0b41000b870501087f20002802c44121010240024041002d00b643450d0041002802b84321070c010b3f002107410041013a00b6434100200741107422073602b8430b200721030240024002400240200741ffff036a41107622023f0022084d0d00200220086b40001a4100210820023f00470d0141002802b84321030b41002108410020033602b84320074100480d0020002001410c6c6a210220074180800441808008200741ffff037122084181f8034922061b6a2008200741ffff077120061b6b20076b2107024041002d00b6430d003f002103410041013a00b6434100200341107422033602b8430b20024180c0006a210220074100480d01200321060240200741076a417871220520036a41ffff036a41107622083f0022044d0d00200820046b40001a20083f00470d0241002802b84321060b4100200620056a3602b8432003417f460d0120002001410c6c6a22014184c0006a2802002206200228020022086a2003460d020240200820014188c0006a22052802002201460d00200620016a2206200628020041808080807871417c20016b20086a72360200200520022802003602002006200628020041ffffffff07713602000b200041c4c1006a2202200228020041016a220236020020002002410c6c6a22004184c0006a200336020020004180c0006a220820073602000b20080f0b02402002280200220820002001410c6c6a22034188c0006a22012802002207460d0020034184c0006a28020020076a2203200328020041808080807871417c20076b20086a72360200200120022802003602002003200328020041ffffffff07713602000b2000200041c4c1006a220728020041016a22033602c0412007200336020041000f0b2002200820076a36020020020b7b01037f024002402000450d0041002802c84222024101480d004188c10021032002410c6c4188c1006a21010340200341046a2802002202450d010240200241046a20004b0d00200220032802006a20004b0d030b2003410c6a22032001490d000b0b0f0b2000417c6a2203200328020041ffffffff07713602000b0300000b0bd901070041040b04c04900000041100b086f6e6572726f72000041200b06656f73696f000041300b406f6e6572726f7220616374696f6e277320617265206f6e6c792076616c69642066726f6d207468652022656f73696f222073797374656d206163636f756e74000041f0000b057265616400004180010b04676574000041e0c2000b566d616c6c6f635f66726f6d5f6672656564207761732064657369676e656420746f206f6e6c792062652063616c6c6564206166746572205f686561702077617320636f6d706c6574656c7920616c6c6f636174656400"}`,
		},
	}

	for _, test := range tests {
		var action SetCode

		b, err := hex.DecodeString(test.in)
		require.NoError(t, err)

		decoder := NewDecoder(b)
		require.NoError(t, decoder.Decode(&action))
		js, err := json.Marshal(action)
		require.NoError(t, err)
		assert.Equal(t, test.jsonTx, string(js))
	}

}
