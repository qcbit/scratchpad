#!/usr/bin/perl

use strict;
#use warnings;
use File::Find;
no warnings 'File::Find';

# This is what I came up with first time reading the challenge
sub findeach_and_remove_world_write_permissions {
    foreach my $filedscr (@_) {
        find(sub {
                my $mode = (stat($filedscr))[2];
                chmod($mode &07775, $filedscr);
                $mode = (stat($filedscr))[2]
                , $File::Find::name if (/$filedscr/) 
            }, 
        ".");
    }
}

sub remove_world_write_permissions {
    my $filedscr = $_[0];
    my $mode = (stat($filedscr))[2];
    chmod($mode &07775, $filedscr);
    $mode = (stat($filedscr))[2]
}

# I think this one more closely solves the challenge
sub remove_world_write_permissions_under {
    foreach (@_) {
        find(sub {
            &remove_world_write_permissions($_),
            $File::Find::name
            },
        $_);
    }
}

