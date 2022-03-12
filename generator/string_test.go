package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func ExampleString() {
	// Streamer uses String generator to generate rune values.
	check.Stream(check.Streamer(
		func(s string) {
			fmt.Printf("%s\n", s)
		},
		generator.String(),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 󎨐뿐򳍖󊷱󻵿󤄃𬟀띛󔐭𪊟򃼨𷺾򀡜񆀕򙏇󀴲򜲹񝛌񭇽򁁏񊻲􅌨󦍅𑖓粩򪱃󒀬𕻯򄤹􌴸𐀞
	// 𣃍𴏯󌄏󦴴񃔳𔺢򏾁񖪍𺕿𫜬񻱖򫽄𙗔󡤨򥕸𔀓󜊹𥮙󴻓񱪋꥗񆽉񂍈󑴭􊟬򋐊񷙐񜁊񝇫񰴤񙚪򀋧񹢰򶝟󥁳􌜶񵌮󿗧򰰀򸢉򳵾򮕮􊆐󴻗򣗭磸񥜞󆠤񭘦󦗃󷅘󔶳񲤾𢙑󱱌򜅂􀨖󠍦矠鰮󰾕򿻙󼾌󪻾
	// 򱁤򾱴󬹱򼄋󱌒󉚜򳴤󦆓񍽥𑋅򎘨򁫞􃺱񵋙򱊞񃏈񄲰򶧆򸹩𩯏񴁊󛃨󌚐񍩄
	// 𲾪𬡔񮞇񀒃𔒤񚳀󭎲𵊠򺋪􆺳񾣽ꃹ񶉑𱶥󶢤𥔌𠰇񇕔򩡯񐸰󊻀𙩶𭣇򤋖􎦩񖈩⦪񷾯􈭓񗡞򉅙닝󟛿񣖭􉵋󶘶񈌌𒒘𿜼󌿚񪵚𭻐򇧹깄󧇢񥢈󉋛堭􇒐󎯬𜂱򗞾󸓑񩡃򼳖񀝳󖪖󫚛🥲𖯶󹸋𦃻񔗰
	// ꦎ񓣵񩮤󴚾򁕳򕮏󒠦󠊙򗂊񋻝􏂌񣒌񘐫񵌨󶴕򪯆񣯏򭵷򨸜򠦜󔓾򉇣󭪘󴨗𮧞󢌧򷹄򑼜󧲭򞁋󁌪񺋐󤥔񸮼񶷬􎶂񋫜񬁀򤱝򡺨𫏐򤫓𼏃𨶛􎃉􁭚󭃋򲼠񓌚󠊸􄡳蕭񨺆񗄒񍏯񯯴񂼂񭆫󓰣򩮍񣼭󒎌
	// 🽸򮇤񒤣󷲪󣛻뾤񩟘󁞘󼌘񳷐먵󃣴𰥢򑊚򂥇􃊡􊓹𑒊򦳘񛿪끔󙗸􀾓𡏡􁧨󏈀񻏶󤪌𾇟􃬍🍃񧯅
	// 󯠙􌨯󜽾󒡻󾽂󛡢򱜃𡫍󼲙򧿯󐿰񧚵󝤳𦹯񮑏𜌫𗒇싓򄘝񲎘󀴼󁁨𒽀򄢸ﴊ񼹡򺩺򈥄򑟍򌨁񷥜󍺕𪻖򓚒󞭏񞫡𹱞𞜃󰂑򽣕񞣽󠧪󊳷򘖗𣛖󁚡𶷲񘺑󠵷򪝶󀺘鱨򉍭񽁝񀑆
	// 򠷒񷙄􅟓򼻼򱐶􇸄򫚙𕎜󤔨򜤙򟺁򋆧􆂵馹󌭅񏰶󧧀񳴃񐿬􅸷𡪱ࡄ򃱵񡮓󂣐􂯨󥻜󘜡񒗊󀕷𧽿򊗼󓭍󧬚򳜏񞬬𧳓𼃼򮭯󴑌񟒳򟍰􏁹񀌖󒔰󒸈񿘥󿨤񐧅𱧮񡼋򎒳􇃾򖖴񷒥򙲨򨾦񛿘𚹏󲜞𨔴򫦝󪏷򞤿򂅛򧢶򪧼󻾂񥏠񋍚򹩔
	// 򼾗򒕪򕅼󦮳򌹭񗈪𮸀򈉢󒜑󥧃ක񥦱򵰢󑦛󬩻󱲴񪐟򣔗𒲋񏑽񎂘򻼋󓐒򷨤񃌶𘁦񰾿󍜟􈻢𧘁򜣴󫵯񲨬񩾎򞉪򀚧񠞵򟳷󲩟
	// 򶐛󼿗􄅅񄺠𳅸񛱎񢦀𪐳𒡉𞑧󨔚󺦘򤨧񃊯󥓀𠅅􎒢᛭򔞿ꢵ򅩲񲡅󡨩򊏁󘼙񗶐󅩅򸩅񃸤񔻥闽󅩑򶮣𼷄񟬍񆴘ᯀ󭱝򘆙熝򘗸󔪴𝡘󽿵򠠭򅱼𑮦񇨡򫞀󀒃𻄀𲅛񨐴󄴾𖝭򣃅򵁙󛎥𫨮󎅜񺌭򾃨󗰕񙃤𒢨󙿻󨍶𬢺򴼅򄎋󲼤�񙡌񼿥𗉘󚿫綤󹡺ꆓ򉿭󦃟󩀾򭥱󶫅񔏶󘽱񁭳󏮍􍇋𬤅
}

func ExampleString_withConstraints() {
	// Streamer uses String generator to generate rune values.
	check.Stream(check.Streamer(
		func(s string) {
			fmt.Printf("%s\n", s)
		},
		// Passing constraint.String to String generator minimal and maximal unicode code point
		// that will be used and minimal and maximal string length.
		// In this example all rune value will be in range [a-z], and string length in range [2, 10]
		generator.String(
			constraints.String{
				Rune: constraints.Rune{
					MinCodePoint: 'a',
					MaxCodePoint: 'z',
				},
				Length: constraints.Length{
					Min: 2,
					Max: 10,
				},
			},
		),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// swqqwrf
	// udadnrii
	// hszmpsi
	// tjjujja
	// dex
	// zyoagn
	// putcbnmwe
	// inytmz
	// ylamx
	// nzmyacimkq
}
