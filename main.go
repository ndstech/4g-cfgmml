package main

import (
	"bytes"
	"database/sql"
	//"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"path/filepath"
	//"strings"
	"time"
	//iconv "github.com/djimenez/iconv-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yogawa/4g-cfgmml/model"
	"golang.org/x/net/html/charset"
)

var (
	i, j             int
	err              error
	connectionString string
	db               *sql.DB
	stmt             *sql.Stmt
	res              sql.Result
	inputFile        string
	start            time.Time
	elapsed          time.Duration
)

func init() {

	start = time.Now()

	conf, _ := os.Open("config.json")
	defer conf.Close()

	jsondecoder := json.NewDecoder(conf)
	configuration := model.Configuration{}
	err := jsondecoder.Decode(&configuration)
	checkErr(err)

	if len(configuration.Password) == 0 {
		connectionString = configuration.Username + "@/" + configuration.Database + "?charset=utf8"
	} else {
		connectionString = configuration.Username + ":" + configuration.Password + "@/" + configuration.Database + "?charset=utf8"
	}

	db, err = sql.Open("mysql", connectionString)
	checkErr(err)

	fmt.Println("Berhasil terhubung ke database")

	err = db.Ping()
	checkErr(err)

	flag.StringVar(&inputFile, "file", "", "XML file")
	flag.Parse()
}

func main() {

	defer db.Close()

	xmlFile, err := os.Open(inputFile)
	checkErr(err)
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var cfgmml model.Cfgmml
	reader := bytes.NewReader(byteValue)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&cfgmml)
	checkErr(err)

	//baseName := strings.TrimRight(filepath.Base(inputFile), ".XML")
	baseName := "N/A"

	for i := 0; i < len(cfgmml.SPECSYNCDATA.CLASSES); i++ {
		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NES) > 0 {
			baseName = cfgmml.SPECSYNCDATA.CLASSES[i].NES[0].ATTRIBUTES.NENAME
			break
		}
	}

	for i := 0; i < len(cfgmml.SPECSYNCDATA.CLASSES); i++ {

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ALMCURCFGS) > 0 {
			insertAlmcurcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ALMCURCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ALMPORTS) > 0 {
			insertAlmport(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ALMPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ANRS) > 0 {
			insertAnr(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ANRS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ANRMEASUREPARAMS) > 0 {
			insertAnrmeasureparam(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ANRMEASUREPARAMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ANTENNAPORTS) > 0 {
			insertAntennaport(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ANTENNAPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].APPCERTS) > 0 {
			insertAppcert(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].APPCERTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].APPLICATIONS) > 0 {
			insertApplication(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].APPLICATIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ASPARAGROUPS) > 0 {
			insertAsparagroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ASPARAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BASEBANDEQMS) > 0 {
			insertBasebandeqm(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].BASEBANDEQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BBPS) > 0 {
			insertBbp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].BBPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BBUFANS) > 0 {
			insertBbufan(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].BBUFANS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BCCHCFGS) > 0 {
			insertBcchcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].BCCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BLINDNCELLOPTS) > 0 {
			insertBlindncellopt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].BLINDNCELLOPTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BRDRESASSIGNMENTS) > 0 {
			insertBrdresassignment(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].BRDRESASSIGNMENTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CABINETS) > 0 {
			insertCabinet(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CABINETS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPS) > 0 {
			insertCagroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPCELLS) > 0 {
			insertCagroupcell(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPSCELLCFGS) > 0 {
			insertCagroupscellcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPSCELLCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CAMGTCFGS) > 0 {
			insertCamgtcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CAMGTCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CASCADEPORTS) > 0 {
			insertCascadeport(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CASCADEPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLS) > 0 {
			insertCell(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLACBARS) > 0 {
			insertCellacbar(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLACBARS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLACCESSES) > 0 {
			insertCellaccess(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLACCESSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLALGOSWITCHES) > 0 {
			insertCellalgoswitch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLALGOSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLBACKOFFS) > 0 {
			insertCellbackoff(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLBACKOFFS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLBFS) > 0 {
			insertCellbf(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLBFS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLBFMIMOPARACFGS) > 0 {
			insertCellbfmimoparacfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLBFMIMOPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCECFGS) > 0 {
			insertCellcecfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCESCHCFGS) > 0 {
			insertCellceschcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCESCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCHPWRCFGS) > 0 {
			insertCellchpwrcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCHPWRCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCOUNTERPARAGROUPS) > 0 {
			insertCellcounterparagroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCOUNTERPARAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCQIADAPTIVECFGS) > 0 {
			insertCellcqiadaptivecfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCQIADAPTIVECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCQIADJALGOS) > 0 {
			insertCellcqiadjalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCQIADJALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCSIRSPARACFGS) > 0 {
			insertCellcsirsparacfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCSIRSPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCSPCPARAS) > 0 {
			insertCellcspcpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCSPCPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDACQCFGS) > 0 {
			insertCelldacqcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDACQCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLCOMPALGOS) > 0 {
			insertCelldlcompalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLCOMPALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLICICS) > 0 {
			insertCelldlicic(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLICICS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDCCHES) > 0 {
			insertCelldlpcpdcch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDCCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDSCHES) > 0 {
			insertCelldlpcpdsch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDSCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDSCHPAS) > 0 {
			insertCelldlpcpdschpa(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDSCHPAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPHICHES) > 0 {
			insertCelldlpcphich(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPHICHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLSCHALGOS) > 0 {
			insertCelldlschalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLSCHALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDRXPARAS) > 0 {
			insertCelldrxpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDRXPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDRXSPECIALPARAS) > 0 {
			insertCelldrxspecialpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDRXSPECIALPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDSSES) > 0 {
			insertCelldss(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDSSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDYNACBARALGOPARAS) > 0 {
			insertCelldynacbaralgopara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDYNACBARALGOPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLEABALGOPARAS) > 0 {
			insertCelleabalgopara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLEABALGOPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLEICICS) > 0 {
			insertCelleicic(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLEICICS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLEMTCALGOS) > 0 {
			insertCellemtcalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLEMTCALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLHOPARACFGS) > 0 {
			insertCellhoparacfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLHOPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLIDPRDUPTS) > 0 {
			insertCellidprdupt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLIDPRDUPTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLIICSPARAS) > 0 {
			insertCelliicspara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLIICSPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLLOWPOWERS) > 0 {
			insertCelllowpower(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLLOWPOWERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLLTEFLEXBWS) > 0 {
			insertCelllteflexbw(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLLTEFLEXBWS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLLTEFLEXBWITFCFGS) > 0 {
			insertCelllteflexbwitfcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLLTEFLEXBWITFCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMBFCSES) > 0 {
			insertCellmbfcs(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMBFCSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMBMSCFGS) > 0 {
			insertCellmbmscfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMBMSCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMCPARAS) > 0 {
			insertCellmcpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMCPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMIMOPARACFGS) > 0 {
			insertCellmimoparacfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMIMOPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBS) > 0 {
			insertCellmlb(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBAUTOGROUPS) > 0 {
			insertCellmlbautogroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBAUTOGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBHOS) > 0 {
			insertCellmlbho(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBHOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBUESELS) > 0 {
			insertCellmlbuesel(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBUESELS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMMALGOS) > 0 {
			insertCellmmalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMMALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMROS) > 0 {
			insertCellmro(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMROS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLNOACCESSALMPARAS) > 0 {
			insertCellnoaccessalmpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLNOACCESSALMPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLOPS) > 0 {
			insertCellop(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLOPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLPCALGOS) > 0 {
			insertCellpcalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLPCALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLPDCCHALGOS) > 0 {
			insertCellpdcchalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLPDCCHALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLPDCCHCECFGS) > 0 {
			insertCellpdcchcecfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLPDCCHCECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLPRBVALMLBS) > 0 {
			insertCellprbvalmlb(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLPRBVALMLBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLPUCCHALGOS) > 0 {
			insertCellpucchalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLPUCCHALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLQCIPARAS) > 0 {
			insertCellqcipara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLQCIPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACHALGOS) > 0 {
			insertCellrachalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACHALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACHCECFGS) > 0 {
			insertCellrachcecfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACHCECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACTHDS) > 0 {
			insertCellracthd(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACTHDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELS) > 0 {
			insertCellresel(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELGERANS) > 0 {
			insertCellreselgeran(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELGERANS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELUTRANS) > 0 {
			insertCellreselutran(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELUTRANS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRFSHUTDOWNS) > 0 {
			insertCellrfshutdown(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRFSHUTDOWNS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRICALGOS) > 0 {
			insertCellricalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRICALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSELS) > 0 {
			insertCellsel(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSELS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSERVICEDIFFCFGS) > 0 {
			insertCellservicediffcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSERVICEDIFFCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSHUTDOWNS) > 0 {
			insertCellshutdown(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSHUTDOWNS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSIMAPS) > 0 {
			insertCellsimap(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSIMAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSRLTES) > 0 {
			insertCellsrlte(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSRLTES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSRSADAPTIVECFGS) > 0 {
			insertCellsrsadaptivecfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSRSADAPTIVECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSTANDARDQCIS) > 0 {
			insertCellstandardqci(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSTANDARDQCIS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLTTIBUNDLINGALGOS) > 0 {
			insertCellttibundlingalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLTTIBUNDLINGALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLUCIONPUSCHPARAS) > 0 {
			insertCellucionpuschpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLUCIONPUSCHPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLUEMEASCONTROLCFGS) > 0 {
			insertCelluemeascontrolcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLUEMEASCONTROLCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULCOMPALGOS) > 0 {
			insertCellulcompalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULCOMPALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICALGOS) > 0 {
			insertCellulicalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICICS) > 0 {
			insertCellulicic(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICICS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICICMCPARAS) > 0 {
			insertCellulicicmcpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICICMCPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULMIMOPARACFGS) > 0 {
			insertCellulmimoparacfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULMIMOPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULPCCOMMS) > 0 {
			insertCellulpccomm(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULPCCOMMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULPCDEDICS) > 0 {
			insertCellulpcdedic(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULPCDEDICS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULSCHALGOS) > 0 {
			insertCellulschalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULSCHALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULUNIFIEDOLCS) > 0 {
			insertCellulunifiedolc(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULUNIFIEDOLCS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLUSPARACFGS) > 0 {
			insertCellusparacfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLUSPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLVMSES) > 0 {
			insertCellvms(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLVMSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLWTTXPARACFGS) > 0 {
			insertCellwttxparacfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLWTTXPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CEPUCCHCFGS) > 0 {
			insertCepucchcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CEPUCCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERACHCFGS) > 0 {
			insertCerachcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CERACHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERTCFGS) > 0 {
			insertCertcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CERTCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERTCHKTSKS) > 0 {
			insertCertchktsk(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CERTCHKTSKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERTDEPLOIES) > 0 {
			insertCertdeploy(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CERTDEPLOIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERTMKS) > 0 {
			insertCertmk(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CERTMKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERTREQS) > 0 {
			insertCertreq(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CERTREQS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CHKS) > 0 {
			insertChk(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CHKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CLKDETECTS) > 0 {
			insertClkdetect(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CLKDETECTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CLZEROBUFFERZONES) > 0 {
			insertClzerobufferzone(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CLZEROBUFFERZONES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORS) > 0 {
			insertCnoperator(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORHOCFGS) > 0 {
			insertCnoperatorhocfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORHOCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORQCIPARAS) > 0 {
			insertCnoperatorqcipara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORQCIPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORSTANDARDQCIS) > 0 {
			insertCnoperatorstandardqci(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORSTANDARDQCIS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORTAS) > 0 {
			insertCnoperatorta(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORTAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].COUNTERCHECKPARAS) > 0 {
			insertCountercheckpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].COUNTERCHECKPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CPBEARERS) > 0 {
			insertCpbearer(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CPBEARERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CPRIPORTS) > 0 {
			insertCpriport(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CPRIPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CPSWITCHES) > 0 {
			insertCpswitch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CPSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CQIADAPTIVECFGS) > 0 {
			insertCqiadaptivecfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CQIADAPTIVECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CRLPOLICIES) > 0 {
			insertCrlpolicy(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CRLPOLICIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKBLINDHOCFGS) > 0 {
			insertCsfallbackblindhocfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKBLINDHOCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKHOS) > 0 {
			insertCsfallbackho(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKHOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKPOLICYCFGS) > 0 {
			insertCsfallbackpolicycfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKPOLICYCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CSPCALGOPARAS) > 0 {
			insertCspcalgopara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].CSPCALGOPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DEVIPS) > 0 {
			insertDevip(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].DEVIPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DHCPRELAYSWITCHES) > 0 {
			insertDhcprelayswitch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].DHCPRELAYSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DHCPSVRIPS) > 0 {
			insertDhcpsvrip(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].DHCPSVRIPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DHCPSWS) > 0 {
			insertDhcpsw(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].DHCPSWS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DIFPRIS) > 0 {
			insertDifpri(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].DIFPRIS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DISTBASEDHOS) > 0 {
			insertDistbasedho(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].DISTBASEDHOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DLFLOWCTRLPARAS) > 0 {
			insertDlflowctrlpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].DLFLOWCTRLPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DRXES) > 0 {
			insertDrx(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].DRXES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DRXPARAGROUPS) > 0 {
			insertDrxparagroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].DRXPARAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DSCPMAPS) > 0 {
			insertDscpmap(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].DSCPMAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].E1T1S) > 0 {
			insertE1t1(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].E1T1S)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].E1T1BEARS) > 0 {
			insertE1t1bear(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].E1T1BEARS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].E1T1BERS) > 0 {
			insertE1t1ber(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].E1T1BERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EMCS) > 0 {
			insertEmc(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EMCS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EMUS) > 0 {
			insertEmu(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EMUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLOPRSVDPARAS) > 0 {
			insertEnbcelloprsvdpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLOPRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLQCIRSVDPARAS) > 0 {
			insertEnbcellqcirsvdpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLQCIRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLRSVDPARAS) > 0 {
			insertEnbcellrsvdpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBCNOPQCIRSVDPARAS) > 0 {
			insertEnbcnopqcirsvdpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBCNOPQCIRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBCNOPRSVDPARAS) > 0 {
			insertEnbcnoprsvdpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBCNOPRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBLICENSEALMTHDS) > 0 {
			insertEnblicensealmthd(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBLICENSEALMTHDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBQCIRSVDPARAS) > 0 {
			insertEnbqcirsvdpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBQCIRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBRSVDPARAS) > 0 {
			insertEnbrsvdpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENERGYCONS) > 0 {
			insertEnergycon(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENERGYCONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBALGOSWITCHES) > 0 {
			insertEnodebalgoswitch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBALGOSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBALMCFGS) > 0 {
			insertEnodebalmcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBALMCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBAUTOPOWEROFFS) > 0 {
			insertEnodebautopoweroff(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBAUTOPOWEROFFS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCHROUTPUTCTRLS) > 0 {
			insertEnodebchroutputctrl(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCHROUTPUTCTRLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCIPHERCAPS) > 0 {
			insertEnodebciphercap(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCIPHERCAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCONNSTATETIMERS) > 0 {
			insertEnodebconnstatetimer(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCONNSTATETIMERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFDDBBRESES) > 0 {
			insertEnodebfddbbres(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFDDBBRESES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFLOWCTRLPARAS) > 0 {
			insertEnodebflowctrlpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFLOWCTRLPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFRAMEOFFSETS) > 0 {
			insertEnodebframeoffset(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFRAMEOFFSETS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFUNCTIONS) > 0 {
			insertEnodebfunction(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFUNCTIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBINTEGRITYCAPS) > 0 {
			insertEnodebintegritycap(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBINTEGRITYCAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBMLBS) > 0 {
			insertEnodebmlb(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBMLBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBNBPARAS) > 0 {
			insertEnodebnbpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBNBPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBRESMODEALGOS) > 0 {
			insertEnodebresmodealgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBRESMODEALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBSHARINGMODES) > 0 {
			insertEnodebsharingmode(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBSHARINGMODES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBSONDBCFGS) > 0 {
			insertEnodebsondbcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBSONDBCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBTDDBBRESES) > 0 {
			insertEnodebtddbbres(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBTDDBBRESES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBTGALGCFGS) > 0 {
			insertEnodebtgalgcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBTGALGCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBUSPARACFGS) > 0 {
			insertEnodebusparacfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBUSPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EPGROUPS) > 0 {
			insertEpgroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EPGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EQUIPMENTS) > 0 {
			insertEquipment(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EQUIPMENTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ETHPORTS) > 0 {
			insertEthport(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ETHPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCELLALMCFGS) > 0 {
			insertEucellalmcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCELLALMCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCELLSECTOREQMS) > 0 {
			insertEucellsectoreqm(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCELLSECTOREQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCOMMCELLSECTOREQMS) > 0 {
			insertEucommcellsectoreqm(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCOMMCELLSECTOREQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHCFGS) > 0 {
			insertEucoschcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHDLCOMPCFGS) > 0 {
			insertEucoschdlcompcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHDLCOMPCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHULICSCFGS) > 0 {
			insertEucoschulicscfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHULICSCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUEPCSECS) > 0 {
			insertEuepcsec(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUEPCSECS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANEXTERNALCELLS) > 0 {
			insertEutranexternalcell(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANEXTERNALCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTERFREQNCELLS) > 0 {
			insertEutraninterfreqncell(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTERFREQNCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTERNFREQS) > 0 {
			insertEutraninternfreq(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTERNFREQS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTRAFREQNCELLS) > 0 {
			insertEutranintrafreqncell(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTRAFREQNCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUULCOSCHCFGS) > 0 {
			insertEuulcoschcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EUULCOSCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EXTENDEDQCIS) > 0 {
			insertExtendedqci(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].EXTENDEDQCIS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].FDDRESMODES) > 0 {
			insertFddresmode(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].FDDRESMODES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].FLTCORRENABLECFGS) > 0 {
			insertFltcorrenablecfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].FLTCORRENABLECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].FMUS) > 0 {
			insertFmu(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].FMUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].FTPCLTS) > 0 {
			insertFtpclt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].FTPCLTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].FTPCLTPORTS) > 0 {
			insertFtpcltport(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].FTPCLTPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSABISCPS) > 0 {
			insertGbtsabiscp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSABISCPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSBBRESES) > 0 {
			insertGbtsbbres(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSBBRESES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSENERGYMGTPARAS) > 0 {
			insertGbtsenergymgtpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSENERGYMGTPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSFUNCTIONS) > 0 {
			insertGbtsfunction(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSFUNCTIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSGLOBALPARAS) > 0 {
			insertGbtsglobalpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSGLOBALPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSPATHES) > 0 {
			insertGbtspath(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSPATHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GERANEXTERNALCELLS) > 0 {
			insertGeranexternalcell(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GERANEXTERNALCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GERANINTERFCFGS) > 0 {
			insertGeraninterfcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GERANINTERFCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GERANNCELLS) > 0 {
			insertGeranncell(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GERANNCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GERANNFREQGROUPS) > 0 {
			insertGerannfreqgroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GERANNFREQGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GERANNFREQGROUPARFCNS) > 0 {
			insertGerannfreqgrouparfcn(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GERANNFREQGROUPARFCNS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOBALPROCSWITCHES) > 0 {
			insertGlobalprocswitch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOBALPROCSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLS) > 0 {
			insertGlocell(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLALGPARAS) > 0 {
			insertGlocellalgpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLALGPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLENERGYMGTPARAS) > 0 {
			insertGlocellenergymgtpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLENERGYMGTPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLOTHPARAS) > 0 {
			insertGlocellothpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLOTHPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLRLALMPARAS) > 0 {
			insertGlocellrlalmpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLRLALMPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLRSVDPARAS) > 0 {
			insertGlocellrsvdpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GPSES) > 0 {
			insertGps(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GPSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GTPUS) > 0 {
			insertGtpu(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GTPUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GTRANSPARAS) > 0 {
			insertGtranspara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GTRANSPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GTRANSPARAGWS) > 0 {
			insertGtransparagw(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GTRANSPARAGWS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GTRXGROUPS) > 0 {
			insertGtrxgroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GTRXGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GTRXGROUPSECTOREQMS) > 0 {
			insertGtrxgroupsectoreqm(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].GTRXGROUPSECTOREQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].HIGHSPDADAPTIONPARAS) > 0 {
			insertHighspdadaptionpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].HIGHSPDADAPTIONPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].HOMEASCOMMS) > 0 {
			insertHomeascomm(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].HOMEASCOMMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].HTCDPAS) > 0 {
			insertHtcdpa(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].HTCDPAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IKECFGS) > 0 {
			insertIkecfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IKECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IMAGRPS) > 0 {
			insertImagrp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IMAGRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IMALNKS) > 0 {
			insertImalnk(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IMALNKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INGCHKTSKS) > 0 {
			insertIngchktsk(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INGCHKTSKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERFREQHOGROUPS) > 0 {
			insertInterfreqhogroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERFREQHOGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATCELLSHUTDOWNS) > 0 {
			insertInterratcellshutdown(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATCELLSHUTDOWNS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCDMA1XRTTGROUPS) > 0 {
			insertInterrathocdma1xrttgroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCDMA1XRTTGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCDMAHRPDGROUPS) > 0 {
			insertInterrathocdmahrpdgroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCDMAHRPDGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCOMMS) > 0 {
			insertInterrathocomm(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCOMMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCOMMGROUPS) > 0 {
			insertInterrathocommgroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCOMMGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOGERANGROUPS) > 0 {
			insertInterrathogerangroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOGERANGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOUTRANGROUPS) > 0 {
			insertInterrathoutrangroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOUTRANGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATPOLICYCFGGROUPS) > 0 {
			insertInterratpolicycfggroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATPOLICYCFGGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTRAFREQHOGROUPS) > 0 {
			insertIntrafreqhogroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INTRAFREQHOGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTRARATHOCOMMS) > 0 {
			insertIntrarathocomm(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].INTRARATHOCOMMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IOPSCFGS) > 0 {
			insertIopscfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IOPSCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IPCLKLNKS) > 0 {
			insertIpclklnk(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IPCLKLNKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IPGUARDS) > 0 {
			insertIpguard(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IPGUARDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IPPATHES) > 0 {
			insertIppath(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IPPATHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IPPMSESSIONS) > 0 {
			insertIppmsession(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IPPMSESSIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IPRTS) > 0 {
			insertIprt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IPRTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IRATNCELLCLASSMGTS) > 0 {
			insertIratncellclassmgt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IRATNCELLCLASSMGTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IUBS) > 0 {
			insertIub(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IUBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IUBCPS) > 0 {
			insertIubcp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].IUBCPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LANSWS) > 0 {
			insertLansw(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LANSWS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LANSWITCHPORTS) > 0 {
			insertLanswitchport(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LANSWITCHPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LICENSECONTROLSTRATEGIES) > 0 {
			insertLicensecontrolstrategy(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LICENSECONTROLSTRATEGIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LICRATIOS) > 0 {
			insertLicratio(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LICRATIOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LINECLKS) > 0 {
			insertLineclk(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LINECLKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTATOMRULES) > 0 {
			insertLioptatomrule(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTATOMRULES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTFEATURES) > 0 {
			insertLioptfeature(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTFEATURES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTFUNCTIONS) > 0 {
			insertLioptfunction(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTFUNCTIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTRULES) > 0 {
			insertLioptrule(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTRULES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTRULEMEMBERS) > 0 {
			insertLioptrulemember(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTRULEMEMBERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LLDPGLOBALS) > 0 {
			insertLldpglobal(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LLDPGLOBALS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LOCALETHPORTS) > 0 {
			insertLocalethport(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LOCALETHPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LOCALIPS) > 0 {
			insertLocalip(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LOCALIPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LOCALIP6S) > 0 {
			insertLocalip6(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LOCALIP6S)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LOCALWAPS) > 0 {
			insertLocalwap(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LOCALWAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LOCATIONS) > 0 {
			insertLocation(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LOCATIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LWAMGTCFGS) > 0 {
			insertLwamgtcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].LWAMGTCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MAINSALARMBINDS) > 0 {
			insertMainsalarmbind(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].MAINSALARMBINDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MANRESALMRPTS) > 0 {
			insertManresalmrpt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].MANRESALMRPTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MBMSPARAS) > 0 {
			insertMbmspara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].MBMSPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MMEFEATURECFGS) > 0 {
			insertMmefeaturecfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].MMEFEATURECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MPTS) > 0 {
			insertMpt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].MPTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MPTRESASSIGNMENTS) > 0 {
			insertMptresassignment(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].MPTRESASSIGNMENTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MROS) > 0 {
			insertMro(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].MROS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NBCELLDLSCHCEALGOS) > 0 {
			insertNbcelldlschcealgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NBCELLDLSCHCEALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NBCELLULSCHCEALGOS) > 0 {
			insertNbcellulschcealgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NBCELLULSCHCEALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NCELLCLASSMGTS) > 0 {
			insertNcellclassmgt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NCELLCLASSMGTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NCELLDLRSRPMEASPARAS) > 0 {
			insertNcelldlrsrpmeaspara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NCELLDLRSRPMEASPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NCELLPARACFGS) > 0 {
			insertNcellparacfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NCELLPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NCELLSRSMEASPARAS) > 0 {
			insertNcellsrsmeaspara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NCELLSRSMEASPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NES) > 0 {
			insertNe(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NEMNTS) > 0 {
			insertNemnt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NEMNTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODES) > 0 {
			insertNode(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBALGPARAS) > 0 {
			insertNodebalgpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBALGPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBBBRESES) > 0 {
			insertNodebbbres(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBBBRESES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBCHRLEVELS) > 0 {
			insertNodebchrlevel(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBCHRLEVELS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBCLSPATIMERS) > 0 {
			insertNodebclspatimer(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBCLSPATIMERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBFUNCTIONS) > 0 {
			insertNodebfunction(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBFUNCTIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBLICENSEALMTHDS) > 0 {
			insertNodeblicensealmthd(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBLICENSEALMTHDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBMULTICELLGRPS) > 0 {
			insertNodebmulticellgrp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBMULTICELLGRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBOPTDYNADJPARAS) > 0 {
			insertNodeboptdynadjpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBOPTDYNADJPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBPOWEROUTAGES) > 0 {
			insertNodebpoweroutage(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBPOWEROUTAGES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRESALLOCRULES) > 0 {
			insertNodebresallocrule(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRESALLOCRULES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRSVDPARAS) > 0 {
			insertNodebrsvdpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRULEACTIONPARAS) > 0 {
			insertNodebruleactionpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRULEACTIONPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBSMTHPWRS) > 0 {
			insertNodebsmthpwr(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBSMTHPWRS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBTRFOVERLOADTHDS) > 0 {
			insertNodebtrfoverloadthd(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBTRFOVERLOADTHDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBUEQOSENHANCES) > 0 {
			insertNodebueqosenhance(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBUEQOSENHANCES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NTPCPS) > 0 {
			insertNtpcp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].NTPCPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].OMCHES) > 0 {
			insertOmch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].OMCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].OPS) > 0 {
			insertOp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].OPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].OUTPORTS) > 0 {
			insertOutport(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].OUTPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PARAAUTOOPTCFGS) > 0 {
			insertParaautooptcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PARAAUTOOPTCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PCCFREQCFGS) > 0 {
			insertPccfreqcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PCCFREQCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PCCHCFGS) > 0 {
			insertPcchcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PCCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PDCPROHCPARAS) > 0 {
			insertPdcprohcpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PDCPROHCPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PDSCHCFGS) > 0 {
			insertPdschcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PDSCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PEERCLKS) > 0 {
			insertPeerclk(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PEERCLKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PEUS) > 0 {
			insertPeu(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PEUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PHICHCFGS) > 0 {
			insertPhichcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PHICHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PHYPORTS) > 0 {
			insertPhyport(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PHYPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PLRTHRESHOLDS) > 0 {
			insertPlrthreshold(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PLRTHRESHOLDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PMTUCFGS) > 0 {
			insertPmtucfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PMTUCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PORTPOLICIES) > 0 {
			insertPortpolicy(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PORTPOLICIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PRI2QUES) > 0 {
			insertPri2que(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PRI2QUES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PRIVATECABANDCOMBS) > 0 {
			insertPrivatecabandcomb(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PRIVATECABANDCOMBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PSUISES) > 0 {
			insertPsuis(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PSUISES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PUCCHCFGS) > 0 {
			insertPucchcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PUCCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PUSCHCFGS) > 0 {
			insertPuschcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PUSCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PUSCHPARAMS) > 0 {
			insertPuschparam(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PUSCHPARAMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PWDPOLICIES) > 0 {
			insertPwdpolicy(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].PWDPOLICIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].QCIPARAS) > 0 {
			insertQcipara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].QCIPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].QOEHOCOMMONCFGS) > 0 {
			insertQoehocommoncfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].QOEHOCOMMONCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RACHCFGS) > 0 {
			insertRachcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RACHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RETS) > 0 {
			insertRet(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RETS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RETDEVICEDATAS) > 0 {
			insertRetdevicedata(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RETDEVICEDATAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RETPORTS) > 0 {
			insertRetport(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RETPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RETSUBUNITS) > 0 {
			insertRetsubunit(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RETSUBUNITS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RFUS) > 0 {
			insertRfu(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RFUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RLCPDCPPARAGROUPS) > 0 {
			insertRlcpdcpparagroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RLCPDCPPARAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RRCCONNSTATETIMERS) > 0 {
			insertRrcconnstatetimer(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RRCCONNSTATETIMERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RRUS) > 0 {
			insertRru(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RRUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RRUCHAINS) > 0 {
			insertRruchain(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RRUCHAINS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RRUJOINTCALPARACFGS) > 0 {
			insertRrujointcalparacfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RRUJOINTCALPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RSCGRPS) > 0 {
			insertRscgrp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RSCGRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RSCGRPALGS) > 0 {
			insertRscgrpalg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RSCGRPALGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RXBRANCHES) > 0 {
			insertRxbranch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].RXBRANCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].S1S) > 0 {
			insertS1(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].S1S)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].S1INTERFACES) > 0 {
			insertS1interface(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].S1INTERFACES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].S1REESTTIMERS) > 0 {
			insertS1reesttimer(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].S1REESTTIMERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SAALLNKS) > 0 {
			insertSaallnk(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SAALLNKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCAPPPARACFGS) > 0 {
			insertScappparacfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SCAPPPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCCFREQCFGS) > 0 {
			insertSccfreqcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SCCFREQCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCPOLICIES) > 0 {
			insertScpolicy(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SCPOLICIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCSERVICEQOSES) > 0 {
			insertScserviceqos(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SCSERVICEQOSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCTPHOSTS) > 0 {
			insertSctphost(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SCTPHOSTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCTPLNKS) > 0 {
			insertSctplnk(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SCTPLNKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCTPPEERS) > 0 {
			insertSctppeer(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SCTPPEERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCTPTEMPLATES) > 0 {
			insertSctptemplate(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SCTPTEMPLATES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SECTORS) > 0 {
			insertSector(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SECTORS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SECTOREQMS) > 0 {
			insertSectoreqm(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SECTOREQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEDIFFSETTINGS) > 0 {
			insertServicediffsetting(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEDIFFSETTINGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIDENTIFYPARAS) > 0 {
			insertServiceidentifypara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIDENTIFYPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIFDLEARFCNGRPS) > 0 {
			insertServiceifdlearfcngrp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIFDLEARFCNGRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIFHOCFGGROUPS) > 0 {
			insertServiceifhocfggroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIFHOCFGGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIRHOCFGGROUPS) > 0 {
			insertServiceirhocfggroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIRHOCFGGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SFPS) > 0 {
			insertSfp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SFPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SIMULOADS) > 0 {
			insertSimuload(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SIMULOADS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SINGLEIPSWITCHES) > 0 {
			insertSingleipswitch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SINGLEIPSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SRBCFGS) > 0 {
			insertSrbcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SRBCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SRBRLCPDCPCFGS) > 0 {
			insertSrbrlcpdcpcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SRBRLCPDCPCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SRCIPRTS) > 0 {
			insertSrciprt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SRCIPRTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SRSADAPTIVECFGS) > 0 {
			insertSrsadaptivecfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SRSADAPTIVECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SRSCFGS) > 0 {
			insertSrscfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SRSCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SSLS) > 0 {
			insertSsl(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SSLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].STANDARDQCIS) > 0 {
			insertStandardqci(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].STANDARDQCIS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SUBRACKS) > 0 {
			insertSubrack(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SUBRACKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SYNCETHES) > 0 {
			insertSynceth(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].SYNCETHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TACALGS) > 0 {
			insertTacalg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TACALGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TASMS) > 0 {
			insertTasm(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TASMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TBDSPINFOS) > 0 {
			insertTbdspinfo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TBDSPINFOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TBLANGNOS) > 0 {
			insertTblangno(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TBLANGNOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TCPACKCTRLALGOS) > 0 {
			insertTcpackctrlalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TCPACKCTRLALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TCPACKLIMITALGS) > 0 {
			insertTcpacklimitalg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TCPACKLIMITALGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TCPMSSCTRLS) > 0 {
			insertTcpmssctrl(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TCPMSSCTRLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TCUS) > 0 {
			insertTcu(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TCUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TDDFRAMEOFFSETS) > 0 {
			insertTddframeoffset(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TDDFRAMEOFFSETS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TDDRESMODESWITCHES) > 0 {
			insertTddresmodeswitch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TDDRESMODESWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TIMEALIGNMENTTIMERS) > 0 {
			insertTimealignmenttimer(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TIMEALIGNMENTTIMERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TIMESRCS) > 0 {
			insertTimesrc(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TIMESRCS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TIMETHRDS) > 0 {
			insertTimethrd(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TIMETHRDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TLDRALGS) > 0 {
			insertTldralg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TLDRALGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TLFRSWITCHES) > 0 {
			insertTlfrswitch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TLFRSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TOLCALGS) > 0 {
			insertTolcalg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TOLCALGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TPEALGOS) > 0 {
			insertTpealgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TPEALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TRPS) > 0 {
			insertTrp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TRUSTCERTS) > 0 {
			insertTrustcert(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TRUSTCERTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TWAMPRESPONDERS) > 0 {
			insertTwampresponder(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TWAMPRESPONDERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TXBRANCHES) > 0 {
			insertTxbranch(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TXBRANCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TYPDRBBSRS) > 0 {
			insertTypdrbbsr(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TYPDRBBSRS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TZS) > 0 {
			insertTz(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].TZS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UDPPINGS) > 0 {
			insertUdpping(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UDPPINGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UDTS) > 0 {
			insertUdt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UDTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UDTPARAGRPS) > 0 {
			insertUdtparagrp(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UDTPARAGRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UECOOPERATIONPARAS) > 0 {
			insertUecooperationpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UECOOPERATIONPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UEIUS) > 0 {
			insertUeiu(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UEIUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UESPECDRXPARAGROUPS) > 0 {
			insertUespecdrxparagroup(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UESPECDRXPARAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UETIMERCONSTS) > 0 {
			insertUetimerconst(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UETIMERCONSTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULCSALGOPARAS) > 0 {
			insertUlcsalgopara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULCSALGOPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULINTERFSUPPRESSCFGS) > 0 {
			insertUlinterfsuppresscfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULINTERFSUPPRESSCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLS) > 0 {
			insertUlocell(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLALGPARAS) > 0 {
			insertUlocellalgpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLALGPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLMACEPARAS) > 0 {
			insertUlocellmacepara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLMACEPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLMACHSPARAS) > 0 {
			insertUlocellmachspara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLMACHSPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLNOACCESSPARAS) > 0 {
			insertUlocellnoaccesspara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLNOACCESSPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLR99ALGPARAS) > 0 {
			insertUlocellr99algpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLR99ALGPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLRSCLMTPARAS) > 0 {
			insertUlocellrsclmtpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLRSCLMTPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLRSVDPARAS) > 0 {
			insertUlocellrsvdpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLSECTOREQMS) > 0 {
			insertUlocellsectoreqm(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLSECTOREQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULZEROBUFFERZONES) > 0 {
			insertUlzerobufferzone(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].ULZEROBUFFERZONES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UPPTSINTERFCFGS) > 0 {
			insertUpptsinterfcfg(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UPPTSINTERFCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USBS) > 0 {
			insertUsb(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].USBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USERPLANEHOSTS) > 0 {
			insertUserplanehost(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].USERPLANEHOSTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USERPLANEPEERS) > 0 {
			insertUserplanepeer(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].USERPLANEPEERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USERPRIORITIES) > 0 {
			insertUserpriority(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].USERPRIORITIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USERQCIPRIORITIES) > 0 {
			insertUserqcipriority(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].USERQCIPRIORITIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USERVPFMPARAS) > 0 {
			insertUservpfmpara(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].USERVPFMPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UTRANEXTERNALCELLS) > 0 {
			insertUtranexternalcell(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UTRANEXTERNALCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UTRANNCELLS) > 0 {
			insertUtranncell(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UTRANNCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UTRANNFREQS) > 0 {
			insertUtrannfreq(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UTRANNFREQS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UTRANRANSHARES) > 0 {
			insertUtranranshare(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].UTRANRANSHARES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].VLANCLASSES) > 0 {
			insertVlanclass(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].VLANCLASSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].VLANMAPS) > 0 {
			insertVlanmap(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].VLANMAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].VQMALGOS) > 0 {
			insertVqmalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].VQMALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].VRFS) > 0 {
			insertVrf(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].VRFS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].WEBLMTS) > 0 {
			insertWeblmt(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].WEBLMTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].WTCPPROXYALGOS) > 0 {
			insertWtcpproxyalgo(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].WTCPPROXYALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].X2S) > 0 {
			insertX2(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].X2S)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].X2INTERFACES) > 0 {
			insertX2interface(baseName, cfgmml.SPECSYNCDATA.CLASSES[i].X2INTERFACES)
		}

	}

	elapsed = time.Since(start)
	fmt.Println("\nProcessing time :", elapsed.Seconds(), "second")

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
}
