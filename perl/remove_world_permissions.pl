#!/usr/bin/perl

use strict;
#use warnings;
use File::Find;
no warnings 'File::Find';

# This is what I came up with first time reading the question
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

# I think this one more closely answers the question.
sub remove_world_write_permissions_under {
    foreach (@_) {
        find(sub {
            &remove_world_write_permissions($_),
            $File::Find::name
            },
        $_);
    }
}

my @my_list = qw/t bad_variables.t cpan-bootstrap.pl Makefile.PL.include/;
my @my_list = qw/author cpan-bootstrap.t release taint-mode.t stackable.t local CPANBootstrapper.pm script pm_to_blib Makefile Makefile.PL/;
#findeach_and_remove_world_write_permissions(@my_list);
my @paths = qw/macvim local-lib-2.000_020/;
#remove_world_write_permissions_under(@paths);
remove_world_write_permissions('remove_world_permissions.pl');

