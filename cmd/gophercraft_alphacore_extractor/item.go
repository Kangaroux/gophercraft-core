package main

import (
	"fmt"
	"reflect"

	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
	"github.com/davecgh/go-spew/spew"
	"github.com/superp00t/etc"
)

var allItems = map[string]*models.ItemTemplate{}

type ItemTemplate struct {
	Entry                     int    `xorm:"'entry'"`
	Class                     uint32 `xorm:"'class'"`
	Subclass                  uint32 `xorm:"'subclass'"`
	Name                      string `xorm:"'name'"`
	DisplayID                 uint32 `xorm:"'display_id'"`
	Quality                   uint8  `xorm:"'quality'"`
	Flags                     uint32 `xorm:"'flags'"`
	BuyCount                  uint8  `xorm:"'buy_count'"`
	BuyPrice                  uint32 `xorm:"'buy_price'"`
	SellPrice                 uint32 `xorm:"'sell_price'"`
	InventoryType             uint8  `xorm:"'inventory_type'"`
	AllowableClass            int32  `xorm:"'allowable_class'"`
	AllowableRace             int32  `xorm:"'allowable_race'"`
	ItemLevel                 uint32 `xorm:"'item_level'"`
	RequiredLevel             uint8  `xorm:"'required_level'"`
	RequiredSkill             uint32 `xorm:"'required_skill'"`
	RequiredSkillRank         uint32 `xorm:"'required_skill_rank'"`
	Requiredspell             uint32 `xorm:"'required_spell'"`
	Requiredhonorrank         uint32 `xorm:"'required_honor_rank'"`
	RequiredCityRank          uint32 `xorm:"'required_city_rank'"`
	RequiredReputationFaction uint32 `xorm:"'required_reputation_faction'"`
	RequiredReputationRank    uint32 `xorm:"'required_reputation_rank'"`
	Maxcount                  uint32 `xorm:"'max_count'"`
	Stackable                 uint32 `xorm:"'stackable'"`
	ContainerSlots            uint8  `xorm:"'container_slots'"`
	// StatsCount                uint8   `xorm:"'StatsCount'"`
	Stat_type1   uint8 `xorm:"'stat_type1'"`
	Stat_value1  int32 `xorm:"'stat_value1'"`
	Stat_type2   uint8 `xorm:"'stat_type2'"`
	Stat_value2  int32 `xorm:"'stat_value2'"`
	Stat_type3   uint8 `xorm:"'stat_type3'"`
	Stat_value3  int32 `xorm:"'stat_value3'"`
	Stat_type4   uint8 `xorm:"'stat_type4'"`
	Stat_value4  int32 `xorm:"'stat_value4'"`
	Stat_type5   uint8 `xorm:"'stat_type5'"`
	Stat_value5  int32 `xorm:"'stat_value5'"`
	Stat_type6   uint8 `xorm:"'stat_type6'"`
	Stat_value6  int32 `xorm:"'stat_value6'"`
	Stat_type7   uint8 `xorm:"'stat_type7'"`
	Stat_value7  int32 `xorm:"'stat_value7'"`
	Stat_type8   uint8 `xorm:"'stat_type8'"`
	Stat_value8  int32 `xorm:"'stat_value8'"`
	Stat_type9   uint8 `xorm:"'stat_type9'"`
	Stat_value9  int32 `xorm:"'stat_value9'"`
	Stat_type10  uint8 `xorm:"'stat_type10'"`
	Stat_value10 int32 `xorm:"'stat_value10'"`
	// ScalingStatDistribution int32   `xorm:"'ScalingStatDistribution'"`
	// ScalingStatValue        int32   `xorm:"'ScalingStatValue'"`
	DMG_min1                float32 `xorm:"'dmg_min1'"`
	DMG_max1                float32 `xorm:"'dmg_max1'"`
	DMG_type1               uint8   `xorm:"'dmg_type1'"`
	DMG_min2                float32 `xorm:"'dmg_min2'"`
	DMG_max2                float32 `xorm:"'dmg_max2'"`
	DMG_type2               uint8   `xorm:"'dmg_type2'"`
	DMG_min3                float32 `xorm:"'dmg_min3'"`
	DMG_max3                float32 `xorm:"'dmg_max3'"`
	DMG_type3               uint8   `xorm:"'dmg_type3'"`
	DMG_min4                float32 `xorm:"'dmg_min4'"`
	DMG_max4                float32 `xorm:"'dmg_max4'"`
	DMG_type4               uint8   `xorm:"'dmg_type4'"`
	DMG_min5                float32 `xorm:"'dmg_min5'"`
	DMG_max5                float32 `xorm:"'dmg_max5'"`
	DMG_type5               uint8   `xorm:"'dmg_type5'"`
	Armor                   int32   `xorm:"'armor'"`
	Holy_res                int32   `xorm:"'holy_res'"`
	Fire_res                int32   `xorm:"'fire_res'"`
	Nature_res              int32   `xorm:"'nature_res'"`
	Frost_res               int32   `xorm:"'frost_res'"`
	Shadow_res              int32   `xorm:"'shadow_res'"`
	Arcane_res              int32   `xorm:"'arcane_res'"`
	Delay                   uint32  `xorm:"'delay'"`
	Ammo_type               uint32  `xorm:"'ammo_type'"`
	RangedModRange          float32 `xorm:"'range_mod'"`
	Spellid_1               uint32  `xorm:"'spellid_1'"`
	Spelltrigger_1          uint32  `xorm:"'spelltrigger_1'"`
	Spellcharges_1          int     `xorm:"'spellcharges_1'"`
	SpellppmRate_1          float32 `xorm:"'spellppmRate_1'"`
	Spellcooldown_1         int32   `xorm:"'spellcooldown_1'"`
	Spellcategory_1         uint32  `xorm:"'spellcategory_1'"`
	Spellcategorycooldown_1 int32   `xorm:"'spellcategorycooldown_1'"`
	Spellid_2               uint32  `xorm:"'spellid_2'"`
	Spelltrigger_2          uint32  `xorm:"'spelltrigger_2'"`
	Spellcharges_2          int32   `xorm:"'spellcharges_2'"`
	SpellppmRate_2          float32 `xorm:"'spellppmRate_2'"`
	Spellcooldown_2         int32   `xorm:"'spellcooldown_2'"`
	Spellcategory_2         uint32  `xorm:"'spellcategory_2'"`
	Spellcategorycooldown_2 int32   `xorm:"'spellcategorycooldown_2'"`
	Spellid_3               uint32  `xorm:"'spellid_3'"`
	Spelltrigger_3          uint32  `xorm:"'spelltrigger_3'"`
	Spellcharges_3          int32   `xorm:"'spellcharges_3'"`
	SpellppmRate_3          float32 `xorm:"'spellppmRate_3'"`
	Spellcooldown_3         int32   `xorm:"'spellcooldown_3'"`
	Spellcategory_3         uint32  `xorm:"'spellcategory_3'"`
	Spellcategorycooldown_3 int32   `xorm:"'spellcategorycooldown_3'"`
	Spellid_4               uint32  `xorm:"'spellid_4'"`
	Spelltrigger_4          uint32  `xorm:"'spelltrigger_4'"`
	Spellcharges_4          int32   `xorm:"'spellcharges_4'"`
	SpellppmRate_4          float32 `xorm:"'spellppmRate_4'"`
	Spellcooldown_4         int32   `xorm:"'spellcooldown_4'"`
	Spellcategory_4         uint32  `xorm:"'spellcategory_4'"`
	Spellcategorycooldown_4 int32   `xorm:"'spellcategorycooldown_4'"`
	Spellid_5               uint32  `xorm:"'spellid_5'"`
	Spelltrigger_5          uint32  `xorm:"'spelltrigger_5'"`
	Spellcharges_5          int32   `xorm:"'spellcharges_5'"`
	SpellppmRate_5          float32 `xorm:"'spellppmRate_5'"`
	Spellcooldown_5         int32   `xorm:"'spellcooldown_5'"`
	Spellcategory_5         uint32  `xorm:"'spellcategory_5'"`
	Spellcategorycooldown_5 int32   `xorm:"'spellcategorycooldown_5'"`
	Bonding                 uint8   `xorm:"'bonding'"`
	Description             string  `xorm:"'description'"`
	PageText                uint32  `xorm:"'page_text'"`
	LanguageID              uint32  `xorm:"'page_language'"`
	PageMaterial            uint32  `xorm:"'page_material'"`
	Startquest              uint32  `xorm:"'start_quest'"`
	Lockid                  uint32  `xorm:"'lock_id'"`
	Material                int32   `xorm:"'material'"`
	Sheath                  uint32  `xorm:"'sheath'"`
	RandomProperty          uint32  `xorm:"'random_property'"`
	// RandomSuffix            uin t32  `xorm:"'RandomSuffix'"`
	Block         uint32 `xorm:"'block'"`
	Itemset       uint32 `xorm:"'set_id'"`
	MaxDurability uint32 `xorm:"'max_durability'"`
	Area          uint32 `xorm:"'area_bound'"`
	Map           int32  `xorm:"'map_bound'"`
	BagFamily     int32  `xorm:"'bag_family'"`
	// TotemCategory           int32   `xorm:"'TotemCategory'"`
	// SocketColor_1           int32   `xorm:"'socketColor_1'"`
	// SocketContent_1         int32   `xorm:"'socketContent_1'"`
	// SocketColor_2           int32   `xorm:"'socketColor_2'"`
	// SocketContent_2         int32   `xorm:"'socketContent_2'"`
	// SocketColor_3           int32   `xorm:"'socketColor_3'"`
	// SocketContent_3         int32   `xorm:"'socketContent_3'"`
	// SocketBonus             int32   `xorm:"'socketBonus'"`
	// GemProperties           int32   `xorm:"'GemProperties'"`
	// RequiredDisenchantSkill int32   `xorm:"'RequiredDisenchantSkill'"`
	// ArmorDamageModifier float32 `xorm:"'ArmorDamageModifier'"`
	// ItemLimitCategory int32  `xorm:"'ItemLimitCategory'"`
	// ScriptName   string `xorm:"'ScriptName'"`
	DisenchantID uint32 `xorm:"'disenchant_id'"`
	FoodType     uint8  `xorm:"'food_type'"`
	MinMoneyLoot uint32 `xorm:"'min_money_loot'"`
	MaxMoneyLoot uint32 `xorm:"'max_money_loot'"`
	Duration     int32  `xorm:"'Duration'"`
	ExtraFlags   uint8  `xorm:"'extra_flags'"`
}

func (it ItemTemplate) GetItemStat(idx int) models.ItemStat {
	st := fmt.Sprintf("Stat_type%d", idx)
	sv := fmt.Sprintf("Stat_value%d", idx)
	val := reflect.ValueOf(it)
	_, ok := val.Type().FieldByName(st)
	if !ok {
		panic("no field for " + st)
	}

	code := val.FieldByName(st).Uint()

	var mod models.ItemMod

	if err := mod.Resolve(vsn.Alpha, uint32(code)); err != nil {
		panic(err)
	}

	return models.ItemStat{
		mod,
		int32(val.FieldByName(sv).Int()),
	}
}

func (it ItemTemplate) GetItemDamage(idx int) models.ItemDamage {
	dt := fmt.Sprintf("DMG_type%d", idx)
	dmn := fmt.Sprintf("DMG_min%d", idx)
	dmx := fmt.Sprintf("DMG_max%d", idx)

	val := reflect.ValueOf(it)
	_, ok := val.Type().FieldByName(dt)
	if !ok {
		panic("no field for " + dt)
	}

	return models.ItemDamage{
		uint8(val.FieldByName(dt).Uint()),
		float32(val.FieldByName(dmn).Float()),
		float32(val.FieldByName(dmx).Float()),
	}
}

func (it ItemTemplate) GetItemSpell(idx int) models.ItemSpell {
	st := fmt.Sprintf("Spellid_%d", idx)
	strigger := fmt.Sprintf("Spelltrigger_%d", idx)
	scharges := fmt.Sprintf("Spellcharges_%d", idx)
	ppm := fmt.Sprintf("SpellppmRate_%d", idx)
	cool := fmt.Sprintf("Spellcooldown_%d", idx)
	scategory := fmt.Sprintf("Spellcategory_%d", idx)
	scategoryCooldown := fmt.Sprintf("Spellcategorycooldown_%d", idx)

	val := reflect.ValueOf(it)
	_, ok := val.Type().FieldByName(st)
	if !ok {
		panic("no field for " + st)
	}

	return models.ItemSpell{
		uint32(val.FieldByName(st).Uint()),
		uint32(val.FieldByName(strigger).Uint()),
		int32(val.FieldByName(scharges).Int()),
		float32(val.FieldByName(ppm).Float()),
		int64(val.FieldByName(cool).Int()),
		uint32(val.FieldByName(scategory).Uint()),
		int64(val.FieldByName(scategoryCooldown).Int()),
	}
}

func extractItems() {
	var itt []ItemTemplate
	err := DB.Find(&itt)
	if err != nil {
		fmt.Println(spew.Sdump(err))
		panic(err)
	}

	fl := openFile("DB/ItemTemplate.txt")

	printTimestamp(fl)

	wr := openTextWriter(fl)

	for _, t := range itt {
		if t.Entry == 0 {
			continue
		}

		var stats []models.ItemStat

		for i := 0; i < 10; i++ {
			st := t.GetItemStat(i + 1)
			if st.Type != 0 {
				stats = append(stats, st)
			}
		}

		var dmg []models.ItemDamage
		for i := 0; i < 5; i++ {
			d := t.GetItemDamage(i + 1)
			if d.Type != 0 {
				dmg = append(dmg, d)
			}
		}

		var spell []models.ItemSpell
		for i := 0; i < 5; i++ {
			sp := t.GetItemSpell(i + 1)
			if sp.ID != 0 {
				spell = append(spell, sp)
			}
		}

		e := etc.NewBuffer()
		e.WriteUint64(uint64(t.Flags))

		var flags update.ItemFlags

		if t.Flags != 0 {
			flags.Bitmask = update.Bitmask{
				t.Flags,
			}
		}

		witem := &models.ItemTemplate{
			ID:                        fmt.Sprintf("it:%d", t.Entry),
			Class:                     t.Class,
			Subclass:                  t.Subclass,
			Name:                      i18n.GetEnglish(t.Name),
			DisplayID:                 t.DisplayID,
			Quality:                   models.ItemQuality(t.Quality),
			Flags:                     flags,
			BuyCount:                  t.BuyCount,
			BuyPrice:                  models.Money(t.BuyPrice),
			SellPrice:                 models.Money(t.SellPrice),
			InventoryType:             models.InventoryType(t.InventoryType),
			AllowableClass:            models.ClassMask(t.AllowableClass),
			AllowableRace:             models.RaceMask(t.AllowableRace),
			ItemLevel:                 t.ItemLevel,
			RequiredLevel:             t.RequiredLevel,
			RequiredSkill:             t.RequiredSkill,
			RequiredSkillRank:         t.RequiredSkillRank,
			RequiredSpell:             t.Requiredspell,
			RequiredHonorRank:         t.Requiredhonorrank,
			RequiredCityRank:          t.RequiredCityRank,
			RequiredReputationFaction: t.RequiredReputationFaction,
			MaxCount:                  t.Maxcount,
			Stackable:                 t.Stackable,
			ContainerSlots:            t.ContainerSlots,
			Stats:                     stats,
			Damage:                    dmg,
			Armor:                     uint32(t.Armor),
			HolyRes:                   uint32(t.Holy_res),
			FireRes:                   uint32(t.Fire_res),
			NatureRes:                 uint32(t.Nature_res),
			FrostRes:                  uint32(t.Frost_res),
			ShadowRes:                 uint32(t.Shadow_res),
			ArcaneRes:                 uint32(t.Arcane_res),
			Delay:                     t.Delay,
			AmmoType:                  t.Ammo_type,
			RangedModRange:            t.RangedModRange,
			Spells:                    spell,
			Bonding:                   models.ItemBind(t.Bonding),
			Description:               i18n.GetEnglish(t.Description),
			PageText:                  t.PageText,
			LanguageID:                t.LanguageID,
			PageMaterial:              t.PageMaterial,
			StartQuest:                t.Startquest,
			LockID:                    t.Lockid,
			Material:                  t.Material,
			Sheath:                    t.Sheath,
			RandomProperty:            t.RandomProperty,
			Block:                     t.Block,
			Itemset:                   t.Itemset,
			MaxDurability:             t.MaxDurability,
			Area:                      t.Area,
			Map:                       t.Map,
			BagFamily:                 t.BagFamily,
			// ScriptName:                t.ScriptName,
			DisenchantID: t.DisenchantID,
			FoodType:     t.FoodType,
			MinMoneyLoot: models.Money(t.MinMoneyLoot),
			MaxMoneyLoot: models.Money(t.MaxMoneyLoot),
			Duration:     t.Duration,
			ExtraFlags:   t.ExtraFlags,
		}
		if err := wr.Encode(witem); err != nil {
			panic(err)
		}
		allItems[witem.ID] = witem
	}

	fl.Close()

	itt = nil
}
