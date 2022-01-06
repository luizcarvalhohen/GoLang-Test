package main

import (
    "testing"

)

func TestLogic(t *testing.T){
    // testValidity
    // success
    test_validity_success_got := testValidity("23-ab-11-caba-56-haha")
    test_validity_success_want := true

    if test_validity_success_got != test_validity_success_want {
        t.Errorf("got %t, wanted %t", test_validity_success_got, test_validity_success_want)
    }

    // fail first argument before - is not an int
    test_validity_first_argument_got := testValidity("aa-ab-11-caba-56-haha")
    test_validity_first_argument_want := false

    if test_validity_first_argument_got != test_validity_first_argument_want {
        t.Errorf("got %t, wanted %t", test_validity_first_argument_got, test_validity_first_argument_want)
    }
    // fail second argument after - is not string
    test_validity_second_argument_got := testValidity("22-25-11-caba-56-haha")
    test_validity_second_argument_want := false

    if test_validity_second_argument_got != test_validity_second_argument_want {
        t.Errorf("got %t, wanted %t", test_validity_second_argument_got, test_validity_second_argument_want)
    }
    // fail their is an empty value between -
    test_validity_empty_got := testValidity("22-aa--caba-56-haha")
    test_validity_empty_want := false

    if test_validity_empty_got != test_validity_empty_want {
        t.Errorf("got %t, wanted %t", test_validity_empty_got, test_validity_empty_want)
    }

    // averageNumber
    // success
    average_number_success_got := averageNumber("23-ab-11-caba-56-haha")
    average_number_success_want := int64(30)

    if average_number_success_got != average_number_success_want {
        t.Errorf("got %d, wanted %d", average_number_success_got, average_number_success_want)
    }
    // fail
    average_number_fail_got := int64(40)
    average_number_fail_want :=  averageNumber("23-ab-50-caba-56-haha")

    if average_number_fail_got != average_number_fail_want {
        t.Errorf("got %d, wanted %d", average_number_fail_got, average_number_fail_want)
    }

    // wholeStory
    whole_story_success_got := wholeStory("23-ab-11-caba-56-haha")
    whole_story_success_want := "ab caba haha"

    if whole_story_success_got != whole_story_success_want {
        t.Errorf("got %s, wanted %s", whole_story_success_got, whole_story_success_want)
    }
    // fail
    whole_story_fail_got := "ab caba haha hjjh"
    whole_story_fail_want := wholeStory("23-ab-11-caba-56-haha")

    if whole_story_fail_got != whole_story_fail_want {
        t.Errorf("got %s, wanted %s", whole_story_fail_got, whole_story_fail_want)
    }

}