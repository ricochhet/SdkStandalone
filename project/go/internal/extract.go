/*
 * PortableBuildTools
 * Copyright (C) 2024 PortableBuildTools contributors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package internal

import (
	"errors"

	aflag "github.com/ricochhet/portablebuildtools/flag"
)

var errMSIExtractMissing = errors.New("MSIExtract tool was not found")

func ExtractMSI(flags *aflag.Flags, args ...string) error {
	if exists, err := aflag.IsFile("./MSIExtract.exe"); err != nil || !exists {
		return errMSIExtractMissing
	}

	if flags.MSIExtractVerbose {
		args = append(args, "-s")
		return Exec("./MSIExtract.exe", args...)
	}

	return Exec("./MSIExtract.exe", args...)
}