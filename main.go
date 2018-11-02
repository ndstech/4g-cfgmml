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
	"path/filepath"
	//"strings"
	"time"
	//iconv "github.com/djimenez/iconv-go"
	//"github.com/fatih/color"
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
	conf             *os.File
)

func init() {

	start = time.Now()

	working_dir, err := os.Getwd()
	checkErr(err)
	exe_path, _ := os.Executable()
	binary_dir, err := filepath.Abs(filepath.Dir(exe_path))
	checkErr(err)

	if _, err := os.Stat(working_dir + "/config.json"); !os.IsNotExist(err) {
		fmt.Println("[+] Found onfiguration file : " + working_dir + "/config.json")
		conf, _ = os.Open(working_dir + "/config.json")
		defer conf.Close()
	} else if _, err := os.Stat(binary_dir + "/config.json"); !os.IsNotExist(err) {
		fmt.Println("[+] Found onfiguration file : " + binary_dir + "/config.json")
		conf, _ = os.Open(binary_dir + "/config.json")
		defer conf.Close()
	} else {
		fmt.Println("[+] Configuration file doesn't exists!")
		//color.Red("[+] Working directory : " + working_dir)
		//color.Red("[+] Binary directory : " + binary_dir)
		os.Exit(3)
	}

	//time.Sleep(time.Duration(2) * time.Second)

	jsondecoder := json.NewDecoder(conf)
	configuration := model.Configuration{}
	err = jsondecoder.Decode(&configuration)
	checkErr(err)

	if len(configuration.Password) == 0 {
		connectionString = configuration.Username + "@/" + configuration.Database + "?charset=utf8"
	} else {
		connectionString = configuration.Username + ":" + configuration.Password + "@/" + configuration.Database + "?charset=utf8"
	}

	db, err = sql.Open("mysql", connectionString)
	checkErr(err)

	fmt.Println("[+] Connected to the database")

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

	//neName := strings.TrimRight(filepath.Base(inputFile), ".XML")
	neName := "N/A"
	eNodeBId := "N/A"
	foundNeName, foundENodeBId := false, false

	for i := 0; i < len(cfgmml.SPECSYNCDATA.CLASSES); i++ {

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NES) > 0 && !foundNeName {
			neName = cfgmml.SPECSYNCDATA.CLASSES[i].NES[0].ATTRIBUTES.NENAME
			foundNeName = true
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFUNCTIONS) > 0 && !foundENodeBId {
			eNodeBId = cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFUNCTIONS[0].ATTRIBUTES.ENodeBId
			foundENodeBId = true
		}

		if foundNeName && foundENodeBId {
			break
		}

	}

	for i := 0; i < len(cfgmml.SPECSYNCDATA.CLASSES); i++ {

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ALMCURCFGS) > 0 {
			insertAlmcurcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ALMCURCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ALMPORTS) > 0 {
			insertAlmport(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ALMPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ANRS) > 0 {
			insertAnr(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ANRS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ANRMEASUREPARAMS) > 0 {
			insertAnrmeasureparam(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ANRMEASUREPARAMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ANTENNAPORTS) > 0 {
			insertAntennaport(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ANTENNAPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].APPCERTS) > 0 {
			insertAppcert(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].APPCERTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].APPLICATIONS) > 0 {
			insertApplication(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].APPLICATIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ASPARAGROUPS) > 0 {
			insertAsparagroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ASPARAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BASEBANDEQMS) > 0 {
			insertBasebandeqm(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].BASEBANDEQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BBPS) > 0 {
			insertBbp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].BBPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BBUFANS) > 0 {
			insertBbufan(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].BBUFANS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BCCHCFGS) > 0 {
			insertBcchcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].BCCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BLINDNCELLOPTS) > 0 {
			insertBlindncellopt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].BLINDNCELLOPTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].BRDRESASSIGNMENTS) > 0 {
			insertBrdresassignment(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].BRDRESASSIGNMENTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CABINETS) > 0 {
			insertCabinet(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CABINETS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPS) > 0 {
			insertCagroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPCELLS) > 0 {
			insertCagroupcell(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPSCELLCFGS) > 0 {
			insertCagroupscellcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CAGROUPSCELLCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CAMGTCFGS) > 0 {
			insertCamgtcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CAMGTCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CASCADEPORTS) > 0 {
			insertCascadeport(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CASCADEPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLS) > 0 {
			insertCell(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLACBARS) > 0 {
			insertCellacbar(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLACBARS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLACCESSES) > 0 {
			insertCellaccess(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLACCESSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLALGOSWITCHES) > 0 {
			insertCellalgoswitch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLALGOSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLBACKOFFS) > 0 {
			insertCellbackoff(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLBACKOFFS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLBFS) > 0 {
			insertCellbf(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLBFS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLBFMIMOPARACFGS) > 0 {
			insertCellbfmimoparacfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLBFMIMOPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCECFGS) > 0 {
			insertCellcecfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCESCHCFGS) > 0 {
			insertCellceschcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCESCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCHPWRCFGS) > 0 {
			insertCellchpwrcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCHPWRCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCOUNTERPARAGROUPS) > 0 {
			insertCellcounterparagroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCOUNTERPARAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCQIADAPTIVECFGS) > 0 {
			insertCellcqiadaptivecfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCQIADAPTIVECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCQIADJALGOS) > 0 {
			insertCellcqiadjalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCQIADJALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCSIRSPARACFGS) > 0 {
			insertCellcsirsparacfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCSIRSPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLCSPCPARAS) > 0 {
			insertCellcspcpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLCSPCPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDACQCFGS) > 0 {
			insertCelldacqcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDACQCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLCOMPALGOS) > 0 {
			insertCelldlcompalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLCOMPALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLICICS) > 0 {
			insertCelldlicic(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLICICS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDCCHES) > 0 {
			insertCelldlpcpdcch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDCCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDSCHES) > 0 {
			insertCelldlpcpdsch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDSCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDSCHPAS) > 0 {
			insertCelldlpcpdschpa(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPDSCHPAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPHICHES) > 0 {
			insertCelldlpcphich(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLPCPHICHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLSCHALGOS) > 0 {
			insertCelldlschalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDLSCHALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDRXPARAS) > 0 {
			insertCelldrxpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDRXPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDRXSPECIALPARAS) > 0 {
			insertCelldrxspecialpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDRXSPECIALPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDSSES) > 0 {
			insertCelldss(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDSSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLDYNACBARALGOPARAS) > 0 {
			insertCelldynacbaralgopara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLDYNACBARALGOPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLEABALGOPARAS) > 0 {
			insertCelleabalgopara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLEABALGOPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLEICICS) > 0 {
			insertCelleicic(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLEICICS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLEMTCALGOS) > 0 {
			insertCellemtcalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLEMTCALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLHOPARACFGS) > 0 {
			insertCellhoparacfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLHOPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLIDPRDUPTS) > 0 {
			insertCellidprdupt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLIDPRDUPTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLIICSPARAS) > 0 {
			insertCelliicspara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLIICSPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLLOWPOWERS) > 0 {
			insertCelllowpower(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLLOWPOWERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLLTEFLEXBWS) > 0 {
			insertCelllteflexbw(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLLTEFLEXBWS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLLTEFLEXBWITFCFGS) > 0 {
			insertCelllteflexbwitfcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLLTEFLEXBWITFCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMBFCSES) > 0 {
			insertCellmbfcs(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMBFCSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMBMSCFGS) > 0 {
			insertCellmbmscfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMBMSCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMCPARAS) > 0 {
			insertCellmcpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMCPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMIMOPARACFGS) > 0 {
			insertCellmimoparacfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMIMOPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBS) > 0 {
			insertCellmlb(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBAUTOGROUPS) > 0 {
			insertCellmlbautogroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBAUTOGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBHOS) > 0 {
			insertCellmlbho(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBHOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBUESELS) > 0 {
			insertCellmlbuesel(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMLBUESELS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMMALGOS) > 0 {
			insertCellmmalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMMALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLMROS) > 0 {
			insertCellmro(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLMROS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLNOACCESSALMPARAS) > 0 {
			insertCellnoaccessalmpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLNOACCESSALMPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLOPS) > 0 {
			insertCellop(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLOPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLPCALGOS) > 0 {
			insertCellpcalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLPCALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLPDCCHALGOS) > 0 {
			insertCellpdcchalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLPDCCHALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLPDCCHCECFGS) > 0 {
			insertCellpdcchcecfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLPDCCHCECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLPRBVALMLBS) > 0 {
			insertCellprbvalmlb(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLPRBVALMLBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLPUCCHALGOS) > 0 {
			insertCellpucchalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLPUCCHALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLQCIPARAS) > 0 {
			insertCellqcipara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLQCIPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACHALGOS) > 0 {
			insertCellrachalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACHALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACHCECFGS) > 0 {
			insertCellrachcecfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACHCECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACTHDS) > 0 {
			insertCellracthd(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRACTHDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELS) > 0 {
			insertCellresel(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELGERANS) > 0 {
			insertCellreselgeran(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELGERANS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELUTRANS) > 0 {
			insertCellreselutran(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRESELUTRANS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRFSHUTDOWNS) > 0 {
			insertCellrfshutdown(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRFSHUTDOWNS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLRICALGOS) > 0 {
			insertCellricalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLRICALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSELS) > 0 {
			insertCellsel(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSELS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSERVICEDIFFCFGS) > 0 {
			insertCellservicediffcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSERVICEDIFFCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSHUTDOWNS) > 0 {
			insertCellshutdown(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSHUTDOWNS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSIMAPS) > 0 {
			insertCellsimap(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSIMAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSRLTES) > 0 {
			insertCellsrlte(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSRLTES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSRSADAPTIVECFGS) > 0 {
			insertCellsrsadaptivecfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSRSADAPTIVECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLSTANDARDQCIS) > 0 {
			insertCellstandardqci(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLSTANDARDQCIS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLTTIBUNDLINGALGOS) > 0 {
			insertCellttibundlingalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLTTIBUNDLINGALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLUCIONPUSCHPARAS) > 0 {
			insertCellucionpuschpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLUCIONPUSCHPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLUEMEASCONTROLCFGS) > 0 {
			insertCelluemeascontrolcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLUEMEASCONTROLCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULCOMPALGOS) > 0 {
			insertCellulcompalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULCOMPALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICALGOS) > 0 {
			insertCellulicalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICICS) > 0 {
			insertCellulicic(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICICS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICICMCPARAS) > 0 {
			insertCellulicicmcpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULICICMCPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULMIMOPARACFGS) > 0 {
			insertCellulmimoparacfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULMIMOPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULPCCOMMS) > 0 {
			insertCellulpccomm(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULPCCOMMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULPCDEDICS) > 0 {
			insertCellulpcdedic(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULPCDEDICS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULSCHALGOS) > 0 {
			insertCellulschalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULSCHALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLULUNIFIEDOLCS) > 0 {
			insertCellulunifiedolc(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLULUNIFIEDOLCS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLUSPARACFGS) > 0 {
			insertCellusparacfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLUSPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLVMSES) > 0 {
			insertCellvms(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLVMSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CELLWTTXPARACFGS) > 0 {
			insertCellwttxparacfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CELLWTTXPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CEPUCCHCFGS) > 0 {
			insertCepucchcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CEPUCCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERACHCFGS) > 0 {
			insertCerachcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CERACHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERTCFGS) > 0 {
			insertCertcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CERTCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERTCHKTSKS) > 0 {
			insertCertchktsk(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CERTCHKTSKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERTDEPLOIES) > 0 {
			insertCertdeploy(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CERTDEPLOIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERTMKS) > 0 {
			insertCertmk(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CERTMKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CERTREQS) > 0 {
			insertCertreq(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CERTREQS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CHKS) > 0 {
			insertChk(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CHKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CLKDETECTS) > 0 {
			insertClkdetect(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CLKDETECTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CLZEROBUFFERZONES) > 0 {
			insertClzerobufferzone(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CLZEROBUFFERZONES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORS) > 0 {
			insertCnoperator(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORHOCFGS) > 0 {
			insertCnoperatorhocfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORHOCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORQCIPARAS) > 0 {
			insertCnoperatorqcipara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORQCIPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORSTANDARDQCIS) > 0 {
			insertCnoperatorstandardqci(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORSTANDARDQCIS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORTAS) > 0 {
			insertCnoperatorta(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CNOPERATORTAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].COUNTERCHECKPARAS) > 0 {
			insertCountercheckpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].COUNTERCHECKPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CPBEARERS) > 0 {
			insertCpbearer(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CPBEARERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CPRIPORTS) > 0 {
			insertCpriport(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CPRIPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CPSWITCHES) > 0 {
			insertCpswitch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CPSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CQIADAPTIVECFGS) > 0 {
			insertCqiadaptivecfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CQIADAPTIVECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CRLPOLICIES) > 0 {
			insertCrlpolicy(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CRLPOLICIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKBLINDHOCFGS) > 0 {
			insertCsfallbackblindhocfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKBLINDHOCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKHOS) > 0 {
			insertCsfallbackho(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKHOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKPOLICYCFGS) > 0 {
			insertCsfallbackpolicycfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CSFALLBACKPOLICYCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].CSPCALGOPARAS) > 0 {
			insertCspcalgopara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].CSPCALGOPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DEVIPS) > 0 {
			insertDevip(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].DEVIPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DHCPRELAYSWITCHES) > 0 {
			insertDhcprelayswitch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].DHCPRELAYSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DHCPSVRIPS) > 0 {
			insertDhcpsvrip(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].DHCPSVRIPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DHCPSWS) > 0 {
			insertDhcpsw(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].DHCPSWS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DIFPRIS) > 0 {
			insertDifpri(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].DIFPRIS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DISTBASEDHOS) > 0 {
			insertDistbasedho(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].DISTBASEDHOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DLFLOWCTRLPARAS) > 0 {
			insertDlflowctrlpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].DLFLOWCTRLPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DRXES) > 0 {
			insertDrx(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].DRXES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DRXPARAGROUPS) > 0 {
			insertDrxparagroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].DRXPARAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].DSCPMAPS) > 0 {
			insertDscpmap(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].DSCPMAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].E1T1S) > 0 {
			insertE1t1(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].E1T1S)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].E1T1BEARS) > 0 {
			insertE1t1bear(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].E1T1BEARS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].E1T1BERS) > 0 {
			insertE1t1ber(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].E1T1BERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EMCS) > 0 {
			insertEmc(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EMCS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EMUS) > 0 {
			insertEmu(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EMUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLOPRSVDPARAS) > 0 {
			insertEnbcelloprsvdpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLOPRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLQCIRSVDPARAS) > 0 {
			insertEnbcellqcirsvdpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLQCIRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLRSVDPARAS) > 0 {
			insertEnbcellrsvdpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBCELLRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBCNOPQCIRSVDPARAS) > 0 {
			insertEnbcnopqcirsvdpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBCNOPQCIRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBCNOPRSVDPARAS) > 0 {
			insertEnbcnoprsvdpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBCNOPRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBLICENSEALMTHDS) > 0 {
			insertEnblicensealmthd(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBLICENSEALMTHDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBQCIRSVDPARAS) > 0 {
			insertEnbqcirsvdpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBQCIRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENBRSVDPARAS) > 0 {
			insertEnbrsvdpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENBRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENERGYCONS) > 0 {
			insertEnergycon(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENERGYCONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBALGOSWITCHES) > 0 {
			insertEnodebalgoswitch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBALGOSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBALMCFGS) > 0 {
			insertEnodebalmcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBALMCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBAUTOPOWEROFFS) > 0 {
			insertEnodebautopoweroff(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBAUTOPOWEROFFS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCHROUTPUTCTRLS) > 0 {
			insertEnodebchroutputctrl(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCHROUTPUTCTRLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCIPHERCAPS) > 0 {
			insertEnodebciphercap(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCIPHERCAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCONNSTATETIMERS) > 0 {
			insertEnodebconnstatetimer(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBCONNSTATETIMERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFDDBBRESES) > 0 {
			insertEnodebfddbbres(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFDDBBRESES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFLOWCTRLPARAS) > 0 {
			insertEnodebflowctrlpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFLOWCTRLPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFRAMEOFFSETS) > 0 {
			insertEnodebframeoffset(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFRAMEOFFSETS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFUNCTIONS) > 0 {
			insertEnodebfunction(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBFUNCTIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBINTEGRITYCAPS) > 0 {
			insertEnodebintegritycap(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBINTEGRITYCAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBMLBS) > 0 {
			insertEnodebmlb(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBMLBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBNBPARAS) > 0 {
			insertEnodebnbpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBNBPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBRESMODEALGOS) > 0 {
			insertEnodebresmodealgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBRESMODEALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBSHARINGMODES) > 0 {
			insertEnodebsharingmode(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBSHARINGMODES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBSONDBCFGS) > 0 {
			insertEnodebsondbcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBSONDBCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBTDDBBRESES) > 0 {
			insertEnodebtddbbres(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBTDDBBRESES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBTGALGCFGS) > 0 {
			insertEnodebtgalgcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBTGALGCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBUSPARACFGS) > 0 {
			insertEnodebusparacfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ENODEBUSPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EPGROUPS) > 0 {
			insertEpgroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EPGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EQUIPMENTS) > 0 {
			insertEquipment(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EQUIPMENTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ETHPORTS) > 0 {
			insertEthport(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ETHPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCELLALMCFGS) > 0 {
			insertEucellalmcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCELLALMCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCELLSECTOREQMS) > 0 {
			insertEucellsectoreqm(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCELLSECTOREQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCOMMCELLSECTOREQMS) > 0 {
			insertEucommcellsectoreqm(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCOMMCELLSECTOREQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHCFGS) > 0 {
			insertEucoschcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHDLCOMPCFGS) > 0 {
			insertEucoschdlcompcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHDLCOMPCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHULICSCFGS) > 0 {
			insertEucoschulicscfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUCOSCHULICSCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUEPCSECS) > 0 {
			insertEuepcsec(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUEPCSECS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANEXTERNALCELLS) > 0 {
			insertEutranexternalcell(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANEXTERNALCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTERFREQNCELLS) > 0 {
			insertEutraninterfreqncell(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTERFREQNCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTERNFREQS) > 0 {
			insertEutraninternfreq(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTERNFREQS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTRAFREQNCELLS) > 0 {
			insertEutranintrafreqncell(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUTRANINTRAFREQNCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EUULCOSCHCFGS) > 0 {
			insertEuulcoschcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EUULCOSCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].EXTENDEDQCIS) > 0 {
			insertExtendedqci(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].EXTENDEDQCIS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].FDDRESMODES) > 0 {
			insertFddresmode(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].FDDRESMODES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].FLTCORRENABLECFGS) > 0 {
			insertFltcorrenablecfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].FLTCORRENABLECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].FMUS) > 0 {
			insertFmu(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].FMUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].FTPCLTS) > 0 {
			insertFtpclt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].FTPCLTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].FTPCLTPORTS) > 0 {
			insertFtpcltport(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].FTPCLTPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSABISCPS) > 0 {
			insertGbtsabiscp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSABISCPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSBBRESES) > 0 {
			insertGbtsbbres(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSBBRESES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSENERGYMGTPARAS) > 0 {
			insertGbtsenergymgtpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSENERGYMGTPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSFUNCTIONS) > 0 {
			insertGbtsfunction(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSFUNCTIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSGLOBALPARAS) > 0 {
			insertGbtsglobalpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSGLOBALPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GBTSPATHES) > 0 {
			insertGbtspath(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GBTSPATHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GERANEXTERNALCELLS) > 0 {
			insertGeranexternalcell(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GERANEXTERNALCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GERANINTERFCFGS) > 0 {
			insertGeraninterfcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GERANINTERFCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GERANNCELLS) > 0 {
			insertGeranncell(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GERANNCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GERANNFREQGROUPS) > 0 {
			insertGerannfreqgroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GERANNFREQGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GERANNFREQGROUPARFCNS) > 0 {
			insertGerannfreqgrouparfcn(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GERANNFREQGROUPARFCNS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOBALPROCSWITCHES) > 0 {
			insertGlobalprocswitch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOBALPROCSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLS) > 0 {
			insertGlocell(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLALGPARAS) > 0 {
			insertGlocellalgpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLALGPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLENERGYMGTPARAS) > 0 {
			insertGlocellenergymgtpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLENERGYMGTPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLOTHPARAS) > 0 {
			insertGlocellothpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLOTHPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLRLALMPARAS) > 0 {
			insertGlocellrlalmpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLRLALMPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLRSVDPARAS) > 0 {
			insertGlocellrsvdpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GLOCELLRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GPSES) > 0 {
			insertGps(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GPSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GTPUS) > 0 {
			insertGtpu(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GTPUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GTRANSPARAS) > 0 {
			insertGtranspara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GTRANSPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GTRANSPARAGWS) > 0 {
			insertGtransparagw(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GTRANSPARAGWS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GTRXGROUPS) > 0 {
			insertGtrxgroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GTRXGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].GTRXGROUPSECTOREQMS) > 0 {
			insertGtrxgroupsectoreqm(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].GTRXGROUPSECTOREQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].HIGHSPDADAPTIONPARAS) > 0 {
			insertHighspdadaptionpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].HIGHSPDADAPTIONPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].HOMEASCOMMS) > 0 {
			insertHomeascomm(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].HOMEASCOMMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].HTCDPAS) > 0 {
			insertHtcdpa(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].HTCDPAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IKECFGS) > 0 {
			insertIkecfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IKECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IMAGRPS) > 0 {
			insertImagrp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IMAGRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IMALNKS) > 0 {
			insertImalnk(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IMALNKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INGCHKTSKS) > 0 {
			insertIngchktsk(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INGCHKTSKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERFREQHOGROUPS) > 0 {
			insertInterfreqhogroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERFREQHOGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATCELLSHUTDOWNS) > 0 {
			insertInterratcellshutdown(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATCELLSHUTDOWNS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCDMA1XRTTGROUPS) > 0 {
			insertInterrathocdma1xrttgroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCDMA1XRTTGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCDMAHRPDGROUPS) > 0 {
			insertInterrathocdmahrpdgroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCDMAHRPDGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCOMMS) > 0 {
			insertInterrathocomm(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCOMMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCOMMGROUPS) > 0 {
			insertInterrathocommgroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOCOMMGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOGERANGROUPS) > 0 {
			insertInterrathogerangroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOGERANGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOUTRANGROUPS) > 0 {
			insertInterrathoutrangroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATHOUTRANGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATPOLICYCFGGROUPS) > 0 {
			insertInterratpolicycfggroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INTERRATPOLICYCFGGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTRAFREQHOGROUPS) > 0 {
			insertIntrafreqhogroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INTRAFREQHOGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].INTRARATHOCOMMS) > 0 {
			insertIntrarathocomm(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].INTRARATHOCOMMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IOPSCFGS) > 0 {
			insertIopscfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IOPSCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IPCLKLNKS) > 0 {
			insertIpclklnk(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IPCLKLNKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IPGUARDS) > 0 {
			insertIpguard(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IPGUARDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IPPATHES) > 0 {
			insertIppath(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IPPATHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IPPMSESSIONS) > 0 {
			insertIppmsession(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IPPMSESSIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IPRTS) > 0 {
			insertIprt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IPRTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IRATNCELLCLASSMGTS) > 0 {
			insertIratncellclassmgt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IRATNCELLCLASSMGTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IUBS) > 0 {
			insertIub(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IUBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].IUBCPS) > 0 {
			insertIubcp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].IUBCPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LANSWS) > 0 {
			insertLansw(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LANSWS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LANSWITCHPORTS) > 0 {
			insertLanswitchport(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LANSWITCHPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LICENSECONTROLSTRATEGIES) > 0 {
			insertLicensecontrolstrategy(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LICENSECONTROLSTRATEGIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LICRATIOS) > 0 {
			insertLicratio(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LICRATIOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LINECLKS) > 0 {
			insertLineclk(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LINECLKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTATOMRULES) > 0 {
			insertLioptatomrule(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTATOMRULES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTFEATURES) > 0 {
			insertLioptfeature(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTFEATURES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTFUNCTIONS) > 0 {
			insertLioptfunction(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTFUNCTIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTRULES) > 0 {
			insertLioptrule(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTRULES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTRULEMEMBERS) > 0 {
			insertLioptrulemember(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LIOPTRULEMEMBERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LLDPGLOBALS) > 0 {
			insertLldpglobal(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LLDPGLOBALS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LOCALETHPORTS) > 0 {
			insertLocalethport(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LOCALETHPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LOCALIPS) > 0 {
			insertLocalip(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LOCALIPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LOCALIP6S) > 0 {
			insertLocalip6(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LOCALIP6S)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LOCALWAPS) > 0 {
			insertLocalwap(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LOCALWAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LOCATIONS) > 0 {
			insertLocation(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LOCATIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].LWAMGTCFGS) > 0 {
			insertLwamgtcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].LWAMGTCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MAINSALARMBINDS) > 0 {
			insertMainsalarmbind(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].MAINSALARMBINDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MANRESALMRPTS) > 0 {
			insertManresalmrpt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].MANRESALMRPTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MBMSPARAS) > 0 {
			insertMbmspara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].MBMSPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MMEFEATURECFGS) > 0 {
			insertMmefeaturecfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].MMEFEATURECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MPTS) > 0 {
			insertMpt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].MPTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MPTRESASSIGNMENTS) > 0 {
			insertMptresassignment(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].MPTRESASSIGNMENTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].MROS) > 0 {
			insertMro(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].MROS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NBCELLDLSCHCEALGOS) > 0 {
			insertNbcelldlschcealgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NBCELLDLSCHCEALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NBCELLULSCHCEALGOS) > 0 {
			insertNbcellulschcealgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NBCELLULSCHCEALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NCELLCLASSMGTS) > 0 {
			insertNcellclassmgt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NCELLCLASSMGTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NCELLDLRSRPMEASPARAS) > 0 {
			insertNcelldlrsrpmeaspara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NCELLDLRSRPMEASPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NCELLPARACFGS) > 0 {
			insertNcellparacfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NCELLPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NCELLSRSMEASPARAS) > 0 {
			insertNcellsrsmeaspara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NCELLSRSMEASPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NES) > 0 {
			insertNe(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NEMNTS) > 0 {
			insertNemnt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NEMNTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODES) > 0 {
			insertNode(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBALGPARAS) > 0 {
			insertNodebalgpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBALGPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBBBRESES) > 0 {
			insertNodebbbres(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBBBRESES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBCHRLEVELS) > 0 {
			insertNodebchrlevel(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBCHRLEVELS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBCLSPATIMERS) > 0 {
			insertNodebclspatimer(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBCLSPATIMERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBFUNCTIONS) > 0 {
			insertNodebfunction(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBFUNCTIONS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBLICENSEALMTHDS) > 0 {
			insertNodeblicensealmthd(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBLICENSEALMTHDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBMULTICELLGRPS) > 0 {
			insertNodebmulticellgrp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBMULTICELLGRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBOPTDYNADJPARAS) > 0 {
			insertNodeboptdynadjpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBOPTDYNADJPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBPOWEROUTAGES) > 0 {
			insertNodebpoweroutage(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBPOWEROUTAGES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRESALLOCRULES) > 0 {
			insertNodebresallocrule(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRESALLOCRULES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRSVDPARAS) > 0 {
			insertNodebrsvdpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRULEACTIONPARAS) > 0 {
			insertNodebruleactionpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBRULEACTIONPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBSMTHPWRS) > 0 {
			insertNodebsmthpwr(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBSMTHPWRS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBTRFOVERLOADTHDS) > 0 {
			insertNodebtrfoverloadthd(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBTRFOVERLOADTHDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NODEBUEQOSENHANCES) > 0 {
			insertNodebueqosenhance(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NODEBUEQOSENHANCES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].NTPCPS) > 0 {
			insertNtpcp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].NTPCPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].OMCHES) > 0 {
			insertOmch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].OMCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].OPS) > 0 {
			insertOp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].OPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].OUTPORTS) > 0 {
			insertOutport(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].OUTPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PARAAUTOOPTCFGS) > 0 {
			insertParaautooptcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PARAAUTOOPTCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PCCFREQCFGS) > 0 {
			insertPccfreqcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PCCFREQCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PCCHCFGS) > 0 {
			insertPcchcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PCCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PDCPROHCPARAS) > 0 {
			insertPdcprohcpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PDCPROHCPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PDSCHCFGS) > 0 {
			insertPdschcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PDSCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PEERCLKS) > 0 {
			insertPeerclk(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PEERCLKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PEUS) > 0 {
			insertPeu(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PEUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PHICHCFGS) > 0 {
			insertPhichcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PHICHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PHYPORTS) > 0 {
			insertPhyport(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PHYPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PLRTHRESHOLDS) > 0 {
			insertPlrthreshold(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PLRTHRESHOLDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PMTUCFGS) > 0 {
			insertPmtucfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PMTUCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PORTPOLICIES) > 0 {
			insertPortpolicy(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PORTPOLICIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PRI2QUES) > 0 {
			insertPri2que(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PRI2QUES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PRIVATECABANDCOMBS) > 0 {
			insertPrivatecabandcomb(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PRIVATECABANDCOMBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PSUISES) > 0 {
			insertPsuis(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PSUISES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PUCCHCFGS) > 0 {
			insertPucchcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PUCCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PUSCHCFGS) > 0 {
			insertPuschcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PUSCHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PUSCHPARAMS) > 0 {
			insertPuschparam(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PUSCHPARAMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].PWDPOLICIES) > 0 {
			insertPwdpolicy(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].PWDPOLICIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].QCIPARAS) > 0 {
			insertQcipara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].QCIPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].QOEHOCOMMONCFGS) > 0 {
			insertQoehocommoncfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].QOEHOCOMMONCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RACHCFGS) > 0 {
			insertRachcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RACHCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RETS) > 0 {
			insertRet(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RETS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RETDEVICEDATAS) > 0 {
			insertRetdevicedata(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RETDEVICEDATAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RETPORTS) > 0 {
			insertRetport(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RETPORTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RETSUBUNITS) > 0 {
			insertRetsubunit(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RETSUBUNITS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RFUS) > 0 {
			insertRfu(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RFUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RLCPDCPPARAGROUPS) > 0 {
			insertRlcpdcpparagroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RLCPDCPPARAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RRCCONNSTATETIMERS) > 0 {
			insertRrcconnstatetimer(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RRCCONNSTATETIMERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RRUS) > 0 {
			insertRru(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RRUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RRUCHAINS) > 0 {
			insertRruchain(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RRUCHAINS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RRUJOINTCALPARACFGS) > 0 {
			insertRrujointcalparacfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RRUJOINTCALPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RSCGRPS) > 0 {
			insertRscgrp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RSCGRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RSCGRPALGS) > 0 {
			insertRscgrpalg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RSCGRPALGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].RXBRANCHES) > 0 {
			insertRxbranch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].RXBRANCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].S1S) > 0 {
			insertS1(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].S1S)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].S1INTERFACES) > 0 {
			insertS1interface(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].S1INTERFACES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].S1REESTTIMERS) > 0 {
			insertS1reesttimer(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].S1REESTTIMERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SAALLNKS) > 0 {
			insertSaallnk(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SAALLNKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCAPPPARACFGS) > 0 {
			insertScappparacfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SCAPPPARACFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCCFREQCFGS) > 0 {
			insertSccfreqcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SCCFREQCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCPOLICIES) > 0 {
			insertScpolicy(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SCPOLICIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCSERVICEQOSES) > 0 {
			insertScserviceqos(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SCSERVICEQOSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCTPHOSTS) > 0 {
			insertSctphost(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SCTPHOSTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCTPLNKS) > 0 {
			insertSctplnk(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SCTPLNKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCTPPEERS) > 0 {
			insertSctppeer(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SCTPPEERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SCTPTEMPLATES) > 0 {
			insertSctptemplate(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SCTPTEMPLATES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SECTORS) > 0 {
			insertSector(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SECTORS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SECTOREQMS) > 0 {
			insertSectoreqm(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SECTOREQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEDIFFSETTINGS) > 0 {
			insertServicediffsetting(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEDIFFSETTINGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIDENTIFYPARAS) > 0 {
			insertServiceidentifypara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIDENTIFYPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIFDLEARFCNGRPS) > 0 {
			insertServiceifdlearfcngrp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIFDLEARFCNGRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIFHOCFGGROUPS) > 0 {
			insertServiceifhocfggroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIFHOCFGGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIRHOCFGGROUPS) > 0 {
			insertServiceirhocfggroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SERVICEIRHOCFGGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SFPS) > 0 {
			insertSfp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SFPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SIMULOADS) > 0 {
			insertSimuload(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SIMULOADS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SINGLEIPSWITCHES) > 0 {
			insertSingleipswitch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SINGLEIPSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SRBCFGS) > 0 {
			insertSrbcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SRBCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SRBRLCPDCPCFGS) > 0 {
			insertSrbrlcpdcpcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SRBRLCPDCPCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SRCIPRTS) > 0 {
			insertSrciprt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SRCIPRTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SRSADAPTIVECFGS) > 0 {
			insertSrsadaptivecfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SRSADAPTIVECFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SRSCFGS) > 0 {
			insertSrscfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SRSCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SSLS) > 0 {
			insertSsl(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SSLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].STANDARDQCIS) > 0 {
			insertStandardqci(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].STANDARDQCIS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SUBRACKS) > 0 {
			insertSubrack(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SUBRACKS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].SYNCETHES) > 0 {
			insertSynceth(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].SYNCETHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TACALGS) > 0 {
			insertTacalg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TACALGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TASMS) > 0 {
			insertTasm(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TASMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TBDSPINFOS) > 0 {
			insertTbdspinfo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TBDSPINFOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TBLANGNOS) > 0 {
			insertTblangno(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TBLANGNOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TCPACKCTRLALGOS) > 0 {
			insertTcpackctrlalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TCPACKCTRLALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TCPACKLIMITALGS) > 0 {
			insertTcpacklimitalg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TCPACKLIMITALGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TCPMSSCTRLS) > 0 {
			insertTcpmssctrl(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TCPMSSCTRLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TCUS) > 0 {
			insertTcu(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TCUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TDDFRAMEOFFSETS) > 0 {
			insertTddframeoffset(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TDDFRAMEOFFSETS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TDDRESMODESWITCHES) > 0 {
			insertTddresmodeswitch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TDDRESMODESWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TIMEALIGNMENTTIMERS) > 0 {
			insertTimealignmenttimer(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TIMEALIGNMENTTIMERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TIMESRCS) > 0 {
			insertTimesrc(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TIMESRCS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TIMETHRDS) > 0 {
			insertTimethrd(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TIMETHRDS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TLDRALGS) > 0 {
			insertTldralg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TLDRALGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TLFRSWITCHES) > 0 {
			insertTlfrswitch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TLFRSWITCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TOLCALGS) > 0 {
			insertTolcalg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TOLCALGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TPEALGOS) > 0 {
			insertTpealgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TPEALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TRPS) > 0 {
			insertTrp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TRUSTCERTS) > 0 {
			insertTrustcert(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TRUSTCERTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TWAMPRESPONDERS) > 0 {
			insertTwampresponder(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TWAMPRESPONDERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TXBRANCHES) > 0 {
			insertTxbranch(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TXBRANCHES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TYPDRBBSRS) > 0 {
			insertTypdrbbsr(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TYPDRBBSRS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].TZS) > 0 {
			insertTz(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].TZS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UDPPINGS) > 0 {
			insertUdpping(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UDPPINGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UDTS) > 0 {
			insertUdt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UDTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UDTPARAGRPS) > 0 {
			insertUdtparagrp(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UDTPARAGRPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UECOOPERATIONPARAS) > 0 {
			insertUecooperationpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UECOOPERATIONPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UEIUS) > 0 {
			insertUeiu(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UEIUS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UESPECDRXPARAGROUPS) > 0 {
			insertUespecdrxparagroup(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UESPECDRXPARAGROUPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UETIMERCONSTS) > 0 {
			insertUetimerconst(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UETIMERCONSTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULCSALGOPARAS) > 0 {
			insertUlcsalgopara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULCSALGOPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULINTERFSUPPRESSCFGS) > 0 {
			insertUlinterfsuppresscfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULINTERFSUPPRESSCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLS) > 0 {
			insertUlocell(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLALGPARAS) > 0 {
			insertUlocellalgpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLALGPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLMACEPARAS) > 0 {
			insertUlocellmacepara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLMACEPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLMACHSPARAS) > 0 {
			insertUlocellmachspara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLMACHSPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLNOACCESSPARAS) > 0 {
			insertUlocellnoaccesspara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLNOACCESSPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLR99ALGPARAS) > 0 {
			insertUlocellr99algpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLR99ALGPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLRSCLMTPARAS) > 0 {
			insertUlocellrsclmtpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLRSCLMTPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLRSVDPARAS) > 0 {
			insertUlocellrsvdpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLRSVDPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLSECTOREQMS) > 0 {
			insertUlocellsectoreqm(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULOCELLSECTOREQMS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].ULZEROBUFFERZONES) > 0 {
			insertUlzerobufferzone(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].ULZEROBUFFERZONES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UPPTSINTERFCFGS) > 0 {
			insertUpptsinterfcfg(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UPPTSINTERFCFGS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USBS) > 0 {
			insertUsb(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].USBS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USERPLANEHOSTS) > 0 {
			insertUserplanehost(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].USERPLANEHOSTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USERPLANEPEERS) > 0 {
			insertUserplanepeer(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].USERPLANEPEERS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USERPRIORITIES) > 0 {
			insertUserpriority(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].USERPRIORITIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USERQCIPRIORITIES) > 0 {
			insertUserqcipriority(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].USERQCIPRIORITIES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].USERVPFMPARAS) > 0 {
			insertUservpfmpara(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].USERVPFMPARAS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UTRANEXTERNALCELLS) > 0 {
			insertUtranexternalcell(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UTRANEXTERNALCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UTRANNCELLS) > 0 {
			insertUtranncell(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UTRANNCELLS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UTRANNFREQS) > 0 {
			insertUtrannfreq(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UTRANNFREQS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].UTRANRANSHARES) > 0 {
			insertUtranranshare(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].UTRANRANSHARES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].VLANCLASSES) > 0 {
			insertVlanclass(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].VLANCLASSES)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].VLANMAPS) > 0 {
			insertVlanmap(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].VLANMAPS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].VQMALGOS) > 0 {
			insertVqmalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].VQMALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].VRFS) > 0 {
			insertVrf(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].VRFS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].WEBLMTS) > 0 {
			insertWeblmt(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].WEBLMTS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].WTCPPROXYALGOS) > 0 {
			insertWtcpproxyalgo(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].WTCPPROXYALGOS)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].X2S) > 0 {
			insertX2(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].X2S)
		}

		if len(cfgmml.SPECSYNCDATA.CLASSES[i].X2INTERFACES) > 0 {
			insertX2interface(eNodeBId, neName, cfgmml.SPECSYNCDATA.CLASSES[i].X2INTERFACES)
		}

	}

	elapsed = time.Since(start)
	fmt.Printf("\n[+] Processing time : %v second", elapsed.Seconds())

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
}
