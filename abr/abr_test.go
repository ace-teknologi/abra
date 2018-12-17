package abr

import (
	"testing"
)

var MainNames = []*BusinessName{
  {OrganisationName: "Bob"},
  {OrganisationName: "Ted"},
  {OrganisationName: "Harry"},
  {OrganisationName: "Jack"},
}
var MainTradingNames = []*BusinessName{
  {OrganisationName: "Jack"},
  {OrganisationName: "Bob"},
  {OrganisationName: "Ted"},
  {OrganisationName: "Harry"},
}

var OtherTradingNames = []*BusinessName{
  {OrganisationName: "Harry"},
  {OrganisationName: "Jack"},
  {OrganisationName: "Bob"},
  {OrganisationName: "Ted"},
}

var BusinessNames = []*BusinessName{
  {OrganisationName: "Ted"},
  {OrganisationName: "Harry"},
  {OrganisationName: "Jack"},
  {OrganisationName: "Bob"},
}

var HumanNames = []*HumanName{
  {GivenName: "Bob", FamilyName: "Robert"},
  {GivenName: "Ted", FamilyName: "Edward"},
  {GivenName: "Harry", FamilyName: "Henry"},
  {GivenName: "Jack", FamilyName: "John"},
}

var MultipleABNs = []*ABN{
  {IdentifierValue: "99124391073", IsCurrentIndicator: "N"},
  {IdentifierValue: "99124391074", IsCurrentIndicator: "Y"},
  {IdentifierValue: "99124391073", IsCurrentIndicator: "N"},
}

var SingleABN = []*ABN{
  {IdentifierValue: "99124391073", IsCurrentIndicator: "Y"},
}

func TestNameWithAllNames(t *testing.T) {
  businessEntity := BusinessEntity{
    BusinessNames: BusinessNames,
    HumanNames: HumanNames,
    MainNames: MainNames,
    MainTradingNames: MainTradingNames,
    OtherTradingNames: OtherTradingNames,
  }

  if businessEntity.Name() != "Bob" {
    t.Errorf("Expected %v, got %v", "Bob", businessEntity.Name())
  }

	return
}

func TestNameWithNoMainNames(t *testing.T) {
  businessEntity := BusinessEntity{
    BusinessNames: BusinessNames,
    HumanNames: HumanNames,
    MainTradingNames: MainTradingNames,
    OtherTradingNames: OtherTradingNames,
  }

  if businessEntity.Name() != "Jack" {
    t.Errorf("Expected %v, got %v", "Jack", businessEntity.Name())
  }

	return
}

func TestNameWithNoMainNamesOrMainTradingNames(t *testing.T) {
  businessEntity := BusinessEntity{
    BusinessNames: BusinessNames,
    HumanNames: HumanNames,
    OtherTradingNames: OtherTradingNames,
  }

  if businessEntity.Name() != "Harry" {
    t.Errorf("Expected %v, got %v", "Harry", businessEntity.Name())
  }

	return
}

func TestNameWithNoMainNamesOrMainTradingNamesOrOtherTradingNames(t *testing.T) {
  businessEntity := BusinessEntity{
    BusinessNames: BusinessNames,
    HumanNames: HumanNames,
  }

  if businessEntity.Name() != "Ted" {
    t.Errorf("Expected %v, got %v", "Ted", businessEntity.Name())
  }

	return
}

func TestNameWithNoMainNamesOrMainTradingNamesOrOtherTradingNamesOrBusinessNames(t *testing.T) {
  businessEntity := BusinessEntity{
    HumanNames: HumanNames,
  }

  if businessEntity.Name() != "Bob Robert" {
    t.Errorf("Expected %v, got %v", "Bob Robert", businessEntity.Name())
  }

	return
}

func TestABNWithMultiple(t *testing.T) {
  businessEntity := BusinessEntity{
    ABNs: MultipleABNs,
  }

  if businessEntity.ABN() != "99124391074" {
    t.Errorf("Expected %v, got %v", "99124391074", businessEntity.ABN())
  }

  return
}

func TestABNWithSingle(t *testing.T) {
  businessEntity := BusinessEntity{
    ABNs: SingleABN,
  }

  if businessEntity.ABN() != "99124391073" {
    t.Errorf("Expected %v, got %v", "99124391073", businessEntity.ABN())
  }

  return
}
