package dungeon

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/patientplatypus/gorest/config"
)

type ClassFeatures struct {
	BaseHitPoints int                `json:"basehitpoints"`
	Proficiencies ProficiencyTypes   `json:"array"`
	Equipment     []EquipmentChoices `json:"equipment"`
}

type ProficiencyTypes struct {
	Armor           []string `json:"armor"`
	Weapons         []string `json:"weapons"`
	Tools           []string `json:"tools"`
	SavingThrows    []string `json:"savingthrows"`
	Skills          []string `json:"skills"`
	BaseSkillNumber int      `json:"baseskillnumber"`
}

type EquipmentChoices struct {
	ChoiceA []string `json:"choicea"`
	ChoiceB []string `json:"choiceb"`
	ChoiceC []string `json:"choicec"`
}

//1
func BarbarianSet(wg *sync.WaitGroup) {
	barbarian := ClassFeatures{
		BaseHitPoints: 12,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor", "medium armor", "shields"},
			Weapons:         []string{"simple weapons", "martial weapons"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Strength", "Constitution"},
			Skills:          []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"greataxe"},
				ChoiceB: []string{"any martial melee weapon"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"two handaxes"},
				ChoiceB: []string{"any simple weapon"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"an explorer's pack", "four javelins"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(barbarian)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("barbarian", data, wg)
}

//2
func BardSet(wg *sync.WaitGroup) {
	bard := ClassFeatures{
		BaseHitPoints: 8,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor"},
			Weapons:         []string{"simple weapons", "hand crossbows", "longswords", "rapiers", "shortswords"},
			Tools:           []string{"three musical instruments of your choice"},
			SavingThrows:    []string{"Dexterity", "Charisma"},
			Skills:          []string{"Athletics", "Acrobatics", "Sleight of Hand", "Stealth", "Arcana", "History", "Investigation", "Nature", "Religion", "Animal Handling", "Insight", "Medicine", "Perception", "Survival", "Deception", "Intimidation", "Performance", "Persuasion"},
			BaseSkillNumber: 3},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"rapier"},
				ChoiceB: []string{"longsword"},
				ChoiceC: []string{"any simple weapon"}},
			{ChoiceA: []string{"diplomat's pack"},
				ChoiceB: []string{"entertainer's pack"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"lute"},
				ChoiceB: []string{"any musical instrument"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"leather armor", "dagger"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(bard)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("bard", data, wg)
}

//3
func ClericSet(wg *sync.WaitGroup) {
	cleric := ClassFeatures{
		BaseHitPoints: 8,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor", "medium armor", "shields"},
			Weapons:         []string{"simple weapons"},
			Tools:           []string{"None"},
			SavingThrows:    []string{"Wisdom", "Charisma"},
			Skills:          []string{"History", "Insight", "Medicine", "Persuasion", "Religion"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"mace"},
				ChoiceB: []string{"warhammer (if proficient)"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"scale male"},
				ChoiceB: []string{"leather armor"},
				ChoiceC: []string{"chain male (if proficient)"}},
			{ChoiceA: []string{"light crossbow", "20 bolts"},
				ChoiceB: []string{"any simple weapon"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"priest's pack", "explorer's pack"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"shield", "holy symbol"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(cleric)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("cleric", data, wg)
}

//4
func DruidSet(wg *sync.WaitGroup) {
	druid := ClassFeatures{
		BaseHitPoints: 8,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor (no metal)", "medium armor (no metal)", "shields (no metal)"},
			Weapons:         []string{"clubs", "daggers", "darts", "javelins", "maces", "quarterstaffs", "scimitars", "sickles", "slings", "spears"},
			Tools:           []string{"Herbalism kit"},
			SavingThrows:    []string{"Wisdom", "Intelligence"},
			Skills:          []string{"Arcana", "Animal Handling", "Insight", "Medicine", "Nature", "Perception", "Religion", "Survival"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"wooden shield"},
				ChoiceB: []string{"any simple weapon"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"scimitar"},
				ChoiceB: []string{"any simple melee weapon"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"leather armor", "explorer's pack", "druidic focus"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(druid)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("druid", data, wg)
}

//5
func FighterSet(wg *sync.WaitGroup) {
	fighter := ClassFeatures{
		BaseHitPoints: 10,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"all armor", "shields"},
			Weapons:         []string{"simple weapons", "martial weapons"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Strength", "Constitution"},
			Skills:          []string{"Acrobatics", "Animal Handling", "Athletics", "History", "Insight", "Intimidation", "Perception", "Survival"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"chain mail"},
				ChoiceB: []string{"leather armor", "longbow", "20 arrows"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"any martial weapon", "shield"},
				ChoiceB: []string{"two martial weapons"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"light crossbow", "20 bolts"},
				ChoiceB: []string{"two handaxes"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"dungeoneer's pack"},
				ChoiceB: []string{"explorer's pack"},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(fighter)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("fighter", data, wg)
}

//6
func MonkSet(wg *sync.WaitGroup) {
	monk := ClassFeatures{
		BaseHitPoints: 10,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"none"},
			Weapons:         []string{"simple weapons", "shortswords"},
			Tools:           []string{"one type of artisan's tools or one musical instrument"},
			SavingThrows:    []string{"Strength", "Dexterity"},
			Skills:          []string{"Acrobatics", "Athletics", "History", "Insight", "Religion", "Stealth"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"shortsword"},
				ChoiceB: []string{"any simple weapon"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"dungeoneer's pack"},
				ChoiceB: []string{"explorer's pack"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"10 darts"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(monk)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("monk", data, wg)
}

//7
func PaladinSet(wg *sync.WaitGroup) {
	paladin := ClassFeatures{
		BaseHitPoints: 10,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"all armor", "shields"},
			Weapons:         []string{"simple weapons", "martial weapons"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Wisdom", "Charisma"},
			Skills:          []string{"Athletics", "Insight", "Intimidation", "Medicine", "Persuasion", "Religion"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"martial weapon", "shield"},
				ChoiceB: []string{"two martial weapons"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"five javelins"},
				ChoiceB: []string{"any simple melee weapon"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"priest's pack"},
				ChoiceB: []string{"explorer's pack"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"chain mail", "holy symbol"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(paladin)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("paladin", data, wg)
}

//8
func RangerSet(wg *sync.WaitGroup) {
	ranger := ClassFeatures{
		BaseHitPoints: 10,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor", "medium armor", "shields"},
			Weapons:         []string{"simple weapons", "martial weapons"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Strength", "Dexterity"},
			Skills:          []string{"Animal Handling", "Athletics", "Insight", "Investigation", "Nature", "Perception", "Stealth", "Survival"},
			BaseSkillNumber: 3},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"scale mail"},
				ChoiceB: []string{"leather armor"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"two shortswords"},
				ChoiceB: []string{"two simple melee weapon"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"dungeoneer's pack"},
				ChoiceB: []string{"explorer's pack"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"longbow", "20 arrows"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(ranger)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("ranger", data, wg)
}

//9
func RogueSet(wg *sync.WaitGroup) {
	rogue := ClassFeatures{
		BaseHitPoints: 8,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor"},
			Weapons:         []string{"simple weapons", "hand crossbows", "longswords", "rapiers", "shortswords"},
			Tools:           []string{"thieves' tools"},
			SavingThrows:    []string{"Intelligence", "Dexterity"},
			Skills:          []string{"Acrobatics", "Athletics", "Deception", "Insight", "Intimidation", "Investigation", "Perception", "Performance", "Persuasion", "Sleight of Hand", "Stealth"},
			BaseSkillNumber: 4},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"rapier"},
				ChoiceB: []string{"shortsword"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"shortbow", "20 arrows"},
				ChoiceB: []string{"shortsword"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"burglar's pack"},
				ChoiceB: []string{"dungeoneer's pack"},
				ChoiceC: []string{"explorer's pack"}},
			{ChoiceA: []string{"leather armor", "two daggers", "thieves' tools"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(rogue)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("rogue", data, wg)
}

//10
func SorcererSet(wg *sync.WaitGroup) {
	sorcerer := ClassFeatures{
		BaseHitPoints: 6,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"none"},
			Weapons:         []string{"daggers", "darts", "slings", "quarterstaffs", "light crossbows"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Charisma", "Constitution"},
			Skills:          []string{"Arcana", "Deception", "Insight", "Intimidation", "Persuasion", "Religion"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"light crossbow", "20 bolts"},
				ChoiceB: []string{"any simple weapon"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"component pouch"},
				ChoiceB: []string{"arcane focus"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"dungeoneer's pack"},
				ChoiceB: []string{"explorer's pack"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"two daggers"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(sorcerer)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("sorcerer", data, wg)
}

//11
func WarlockSet(wg *sync.WaitGroup) {
	warlock := ClassFeatures{
		BaseHitPoints: 8,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"light armor"},
			Weapons:         []string{"simple weapons"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Charisma", "Wisdom"},
			Skills:          []string{"Arcana", "Deception", "History", "Intimidation", "Investigation", "Nature", "Religion"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"light crossbow", "20 bolts"},
				ChoiceB: []string{"any simple weapon"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"component pouch"},
				ChoiceB: []string{"arcane focus"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"scholar's pack"},
				ChoiceB: []string{"dungeoneer's pack"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"leather armor", "any simple weapon", "two daggers"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(warlock)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("warlock", data, wg)
}

//12
func WizardSet(wg *sync.WaitGroup) {
	wizard := ClassFeatures{
		BaseHitPoints: 6,
		Proficiencies: ProficiencyTypes{
			Armor:           []string{"none"},
			Weapons:         []string{"daggers", "darts", "slings", "quarterstaffs", "light crossbows"},
			Tools:           []string{"none"},
			SavingThrows:    []string{"Intelligence", "Wisdom"},
			Skills:          []string{"Arcana", "History", "Insight", "Investigation", "Investigation", "Medicine", "Religion"},
			BaseSkillNumber: 2},
		Equipment: []EquipmentChoices{
			{ChoiceA: []string{"quarterstaff"},
				ChoiceB: []string{"dagger"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"component pouch"},
				ChoiceB: []string{"arcane focus"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"scholar's pack"},
				ChoiceB: []string{"explorer's pack"},
				ChoiceC: []string{""}},
			{ChoiceA: []string{"spellbook"},
				ChoiceB: []string{""},
				ChoiceC: []string{""}},
		}}
	data, err := json.Marshal(wizard)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}
	CheckClass("wizard", data, wg)
}

func CheckClass(classname string, data []byte, wg *sync.WaitGroup) {
	db := config.Sql_connect()
	if err := db.QueryRow("SELECT classname FROM dungeon_classes WHERE classname=$1", classname).Scan(&classname); err == nil {
		//found name
		log.Print("The chosen class already exists in dungeon_classes!")
		log.Print("Value of all entries in database is: ")
		rows, error := db.Query("Select classname, data from dungeon_classes")
		defer rows.Close()
		for rows.Next() {
			var classname string
			var data []byte
			error = rows.Scan(&classname, &data)
			if error != nil {
				panic(error)
			}
			var obj ClassFeatures
			if err := json.Unmarshal(data, &obj); err != nil {
				panic(err)
			}
			// log.Print("classname: ", classname, " obj ", obj)
			// fmt.Printf("classname: %+v\n", classname, " obj: %+v\n", obj)
			// fmt.Println("classname: ", classname, " obj ", obj)
			fmt.Printf("%+v\n", obj)
		}
		wg.Done()
	} else if err.Error() == "sql: no rows in result set" {
		log.Print("value of err: ", err)
		log.Print("no rows found of class, inserting into db")
		// log.Print("right before value of barbariandata")
		// log.Print("value of barbariandata: ", barbariandata)
		db.QueryRow("INSERT INTO dungeon_classes VALUES ($1, $2);", classname, data)
		wg.Done()
	}
}

func ClassType(w http.ResponseWriter, r *http.Request) {
	log.Print("inside classtype of dungeon")
	db := config.Sql_connect()
	db.Query("CREATE TABLE IF NOT EXISTS dungeon_classes")

	var wg sync.WaitGroup
	wg.Add(12)

	//1
	go BarbarianSet(&wg)
	//2
	go BardSet(&wg)
	//3
	go ClericSet(&wg)
	//4
	go DruidSet(&wg)
	//5
	go FighterSet(&wg)
	//6
	go MonkSet(&wg)
	//7
	go PaladinSet(&wg)
	//8
	go RangerSet(&wg)
	//9
	go RogueSet(&wg)
	//10
	go SorcererSet(&wg)
	//11
	go WarlockSet(&wg)
	//12
	go WizardSet(&wg)

	wg.Wait()
	log.Print("Done")

}
